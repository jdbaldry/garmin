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

func createStepParams(msg *fit.MonitoringMsg) postgresql.CreateStepParams {
	return postgresql.CreateStepParams{
		Ts:              sql.NullTime{Time: msg.Timestamp, Valid: true},
		Distance:        sql.NullFloat64{Float64: msg.GetDistanceScaled(), Valid: true},
		Cycles:          sql.NullFloat64{Float64: msg.GetCyclesScaled(), Valid: true},
		ActiveTime:      sql.NullFloat64{Float64: msg.GetActiveTimeScaled(), Valid: true},
		ActiveCalories:  sql.NullInt32{},
		DurationMin:     sql.NullInt16{},
		ActivityType:    sql.NullInt16{Int16: int16(msg.ActivityType), Valid: true},
		ActivitySubType: sql.NullInt16{Int16: int16(msg.ActivitySubtype), Valid: true},
	}
}

func ingestStep(ctx context.Context, queries *postgresql.Queries, data *fit.File) error {
	stepFile, err := data.MonitoringB()
	if err != nil {
		// Error contains all useful information.
		//nolint:wrapcheck
		return err
	}

	for _, monitoring := range stepFile.Monitorings {
		if monitoring.ActivityType == 1 || monitoring.ActivityType == 6 {
			params := createStepParams(monitoring)

			log.Infof("Creating step with params: %+v", params)

			_, err := queries.CreateStep(ctx, params)
			if err != nil {
				if !errors.Is(err, sql.ErrNoRows) {
					return fmt.Errorf("failed to create step: %w", err)
				}
			}
		}
	}

	return nil
}

func ingestSteps(ctx context.Context, queries *postgresql.Queries) error {
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

		return ingestStep(ctx, queries, data)
	})
}
