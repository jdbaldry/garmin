package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io/fs"
	"path/filepath"
	"time"

	log "github.com/golang/glog"
	"github.com/tormoder/fit"
	"jdb.sh/garmin/postgresql"
)

func createSleepParams(file *fit.SleepFile) postgresql.CreateSleepParams {
	return postgresql.CreateSleepParams{
		StartTs: sql.NullTime{Time: file.Events[0].Timestamp, Valid: true},
		EndTs:   sql.NullTime{Time: file.Events[1].Timestamp, Valid: true},
	}
}

func createSleepRecordParams(msg *fit.SleepLevelMsg, id int64, startTs time.Time) postgresql.CreateSleepRecordParams {
	return postgresql.CreateSleepRecordParams{
		Sleep:              sql.NullInt64{Int64: id, Valid: true},
		StartTs:            sql.NullTime{Time: startTs, Valid: true},
		EndTs:              sql.NullTime{Time: msg.EndTimestamp, Valid: true},
		SleepActivityLevel: sql.NullInt16{Int16: int16(msg.ActivityLevel), Valid: true},
	}
}

func ingestSleep(ctx context.Context, queries *postgresql.Queries, data *fit.File) error {
	sleepFile, err := data.Sleep()
	if err != nil {
		// Error contains all useful information.
		//nolint:wrapcheck
		return err
	}

	params := createSleepParams(sleepFile)
	log.Infof("Creating sleep with params: %+v", params)

	id, err := queries.CreateSleep(ctx, params)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("failed to create sleep: %w", err)
		}
	}

	start := params.StartTs.Time

	fmt.Printf("%#v\n", sleepFile)
	for _, level := range sleepFile.SleepLevels {
		params := createSleepRecordParams(level, id, start)
		log.Infof("Creating sleep record with params: %+v", params)

		_, err := queries.CreateSleepRecord(ctx, params)
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				return fmt.Errorf("failed to create sleep record: %w", err)
			}
		}

		start = params.EndTs.Time
	}

	return nil
}

func ingestSleeps(ctx context.Context, queries *postgresql.Queries) error {
	return filepath.WalkDir(filepath.Join(fitDir, "Sleep"), func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !isFitFile(path) {
			return nil
		}

		data, err := decode(path)
		if err != nil {
			return err
		}

		return ingestSleep(ctx, queries, data)
	})
}
