package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io/fs"
	"path/filepath"

	log "github.com/golang/glog"
	"github.com/tormoder/fit"
	"jdb.sh/garmin/postgresql"
)

func createSleepRecordParams(msg *fit.SleepEventMsg) postgresql.CreateSleepRecordParams {
	return postgresql.CreateSleepRecordParams{
		Sleep:              sql.NullInt64{},
		StartTs:            sql.NullTime{},
		EndTs:              sql.NullTime{},
		SleepActivityLevel: sql.NullInt16{},
	}
}

func ingestSleep(ctx context.Context, queries *postgresql.Queries, data *fit.File) error {
	sleepFile, err := data.Sleep()
	if err != nil {
		return err
	}

	for _, event := range sleepFile.SleepEvents {
		params := createSleepRecordParams(event)
		log.Infof("Creating sleep record with params: %+v", params)

		_, err := queries.CreateSleepRecord(ctx, params)
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				return fmt.Errorf("failed to create sleep record: %w", err)
			}
		}
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
