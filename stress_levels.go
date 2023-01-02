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

func createStressLevelParams(msg *fit.StressLevelMsg) postgresql.CreateStressLevelParams {
	return postgresql.CreateStressLevelParams{
		Ts:    sql.NullTime{Time: msg.StressLevelTime, Valid: true},
		Value: sql.NullInt16{Int16: msg.StressLevelValue, Valid: true},
	}
}

func ingestStressLevel(ctx context.Context, queries *postgresql.Queries, data *fit.File) error {
	monitoringFile, err := data.MonitoringB()
	if err != nil {
		return err
	}

	for _, sl := range monitoringFile.StressLevels {
		params := createStressLevelParams(sl)
		log.Infof("Creating stress level with params: %+v", params)

		_, err := queries.CreateStressLevel(ctx, params)
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				return fmt.Errorf("failed to create stress level: %w", err)
			}
		}
	}

	return nil
}

func ingestStressLevels(ctx context.Context, queries *postgresql.Queries) error {
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

		return ingestStressLevel(ctx, queries, data)
	})
}
