package main

import (
	"context"
	"database/sql"
	"encoding/csv"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"jdb.sh/garmin/postgresql"
)

const (
	epochOffset = 631_065_600
)

var (
	errUnrecognizedLocal     = errors.New("unrecognized local")
	errUnrecognizedEventType = errors.New("unrecognized event type")
	errUnrecognizedSleepType = errors.New("unrecognized sleep type (eighth field)")
	errNoStart               = errors.New("no start event found")

	messages = map[string]int64{
		"0": 0, // File ID
		"1": 1, // File creator
		"2": 2, // Device info
		"3": 3, // Unknown
		"4": 4, // Event
		"5": 5, // Unknown
		"6": 6, // Sleep event (?)
		"7": 7, // Unknown
	}
	sleepActivityLevels = map[string]int64{
		"0": 0, // Unmeasurable
		"1": 1, // Awake
		"2": 2, // Light
		"3": 3, // Deep
		"4": 4, // REM
	}
)

func ingestSleeps(ctx context.Context, queries *postgresql.Queries) error {
	return filepath.WalkDir(filepath.Join(fitDir, "Sleep"), func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if filepath.Ext(path) != ".csv" {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}

		r := csv.NewReader(f)
		// TODO: Understand correct number of fields.
		r.FieldsPerRecord = -1
		records, err := r.ReadAll()
		if err != nil {
			return err
		}

		var start int
		var id int64

		for _, record := range records {
			if record[0] != "Data" {
				continue
			}

			message, ok := messages[record[1]]
			if !ok {
				return errUnrecognizedLocal
			}

			switch message {
			case 4:
				switch record[13] {
				case "0":
					start, _ = strconv.Atoi(record[4])
					start += epochOffset
				case "1":
					end, _ := strconv.Atoi(record[4])
					end += epochOffset

					var err error
					id, err = queries.CreateSleep(ctx, postgresql.CreateSleepParams{
						StartTs: sql.NullTime{Time: time.Unix(int64(start), 0), Valid: true},
						EndTs:   sql.NullTime{Time: time.Unix(int64(end), 0), Valid: true},
					})
					if err != nil {
						if !errors.Is(err, sql.ErrNoRows) {
							return err
						}
					}
				}
			case 6:
				if start == 0 {
					return errNoStart
				}

				if id == 0 {
					continue
				}

				level, ok := sleepActivityLevels[record[7]]
				if !ok {
					return fmt.Errorf("%q: %w", record, errUnrecognizedSleepType)
				}

				end, _ := strconv.Atoi(record[4])
				end += epochOffset

				if _, err := queries.CreateSleepRecord(ctx, postgresql.CreateSleepRecordParams{
					Sleep:              sql.NullInt64{Int64: id, Valid: true},
					StartTs:            sql.NullTime{Time: time.Unix(int64(start), 0), Valid: true},
					EndTs:              sql.NullTime{Time: time.Unix(int64(end), 0), Valid: true},
					SleepActivityLevel: sql.NullInt16{Int16: int16(level), Valid: true},
				}); err != nil {
					return err
				}

				start = end
			}
		}

		return nil
	})
}
