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

// From fit/time.go.
var timeBase = time.Date(1989, time.December, 31, 0, 0, 0, 0, time.UTC)

// From fit/time.go.
func encodeTime(t time.Time) uint32 {
	return uint32(t.Sub(timeBase) / time.Second)
}

// From fit/time.go.
func decodeDateTime(dt uint32) time.Time {
	return timeBase.Add(time.Duration(dt) * time.Second)
}

// createHeartRateParams produces query parameters from a heart rate monitoring message.
// The timestamp_16 field is not a date_time, rather it is a uint16.
// The 16 bit values are used to keep the size of the file small, which is important for files that may contain 24 hours of monitoring information.
// The timestamp_16 values are the value of the previous timestamp field with the upper two bytes masked off plus any time change that occurred between the two monitoring messages.
// https://forums.garmin.com/developer/fit-sdk/f/discussion/311422/fit-timestamp_16-heart-rate---excel
func createHeartRateParams(msg *fit.MonitoringMsg, timestamp time.Time) postgresql.CreateHeartRateParams {
	encoded := encodeTime(timestamp)
	// Mask off lower two bytes.
	encoded &= 0xFFFF0000
	// Add timestamp16.
	encoded += uint32(msg.Timestamp16)

	ts := decodeDateTime(encoded)

	return postgresql.CreateHeartRateParams{
		Ts:    sql.NullTime{Time: ts, Valid: true},
		Value: sql.NullInt16{Int16: int16(msg.HeartRate), Valid: true},
	}
}

func ingestHeartRate(ctx context.Context, queries *postgresql.Queries, data *fit.File) error {
	monitoringFile, err := data.MonitoringB()
	if err != nil {
		return err
	}

	var timestamp time.Time

	for _, mon := range monitoringFile.Monitorings {
		if !fit.IsBaseTime(mon.Timestamp) {
			timestamp = mon.Timestamp
		}

		if mon.HeartRate != 0xFF {
			params := createHeartRateParams(mon, timestamp)
			log.Infof("Creating heart rate with params: %+v", params)

			_, err := queries.CreateHeartRate(ctx, params)
			if err != nil {
				if !errors.Is(err, sql.ErrNoRows) {
					return fmt.Errorf("failed to create stress level: %w", err)
				}
			}
		}
	}

	return nil
}

func ingestHeartRates(ctx context.Context, queries *postgresql.Queries) error {
	return filepath.WalkDir(filepath.Join(fitDir, "Monitor"), func(path string, d fs.DirEntry, err error) error {
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

		return ingestHeartRate(ctx, queries, data)
	})
}
