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

func createActivityParams(msg *fit.ActivityMsg, source string) postgresql.CreateActivityParams {
	return postgresql.CreateActivityParams{
		StartTs: sql.NullTime{
			Time:  msg.Timestamp.Add(-(time.Duration(msg.GetTotalTimerTimeScaled()) * time.Second)),
			Valid: true,
		},
		EndTs:          sql.NullTime{Time: msg.Timestamp, Valid: true},
		TotalTimerTime: sql.NullFloat64{Float64: msg.GetTotalTimerTimeScaled(), Valid: true},
		NumSessions:    sql.NullInt32{Int32: int32(msg.NumSessions), Valid: true},
		Type:           sql.NullInt32{Int32: int32(msg.Type), Valid: true},
		Event:          sql.NullInt16{Int16: int16(msg.Event), Valid: true},
		EventType:      sql.NullInt16{Int16: int16(msg.EventType), Valid: true},
		LocalTs:        sql.NullTime{Time: msg.LocalTimestamp, Valid: true},
		EventGroup:     sql.NullInt16{Int16: int16(msg.EventGroup), Valid: true},
		Source:         sql.NullString{String: source, Valid: true},
	}
}

func createActivitySessionParams(activityID int64, msg *fit.SessionMsg) postgresql.CreateActivitySessionParams {
	return postgresql.CreateActivitySessionParams{
		Activity:         sql.NullInt64{Int64: activityID, Valid: true},
		StartTs:          sql.NullTime{Time: msg.StartTime, Valid: true},
		EndTs:            sql.NullTime{Time: msg.Timestamp, Valid: true},
		Event:            sql.NullInt16{Int16: int16(msg.Event), Valid: true},
		EventType:        sql.NullInt16{Int16: int16(msg.EventType), Valid: true},
		Sport:            sql.NullInt16{Int16: int16(msg.Sport), Valid: true},
		SubSport:         sql.NullInt16{Int16: int16(msg.SubSport), Valid: true},
		TotalElapsedTime: sql.NullFloat64{Float64: msg.GetTotalElapsedTimeScaled(), Valid: true},
		TotalTimerTime:   sql.NullFloat64{Float64: msg.GetTotalTimerTimeScaled(), Valid: true},
		TotalDistance:    sql.NullFloat64{Float64: msg.GetTotalDistanceScaled(), Valid: true},
		TotalCalories:    sql.NullInt16{Int16: int16(msg.TotalCalories), Valid: true},
		AvgSpeed:         sql.NullFloat64{Float64: msg.GetEnhancedAvgSpeedScaled(), Valid: true},
		MaxSpeed:         sql.NullFloat64{Float64: msg.GetEnhancedMaxSpeedScaled(), Valid: true},
		AvgHeartRate:     sql.NullInt16{Int16: int16(msg.AvgHeartRate), Valid: true},
		MaxHeartRate:     sql.NullInt16{Int16: int16(msg.MaxHeartRate), Valid: true},
		AvgVerticalRatio: sql.NullFloat64{Float64: msg.GetAvgVerticalRatioScaled(), Valid: true},
		AvgStanceTime:    sql.NullFloat64{Float64: msg.GetAvgStanceTimeScaled(), Valid: true},
	}
}

func createActivityLapParams(activityID int64, msg *fit.LapMsg) postgresql.CreateActivityLapParams {
	return postgresql.CreateActivityLapParams{
		Activity:         sql.NullInt64{Int64: activityID, Valid: true},
		MessageIndex:     sql.NullInt16{Int16: 0, Valid: false},
		StartTs:          sql.NullTime{Time: msg.StartTime, Valid: true},
		EndTs:            sql.NullTime{Time: msg.Timestamp, Valid: true},
		Event:            sql.NullInt16{Int16: int16(msg.Event), Valid: true},
		EventType:        sql.NullInt16{Int16: int16(msg.EventType), Valid: true},
		Sport:            sql.NullInt16{Int16: int16(msg.Sport), Valid: true},
		SubSport:         sql.NullInt16{Int16: int16(msg.SubSport), Valid: true},
		TotalElapsedTime: sql.NullFloat64{Float64: msg.GetTotalElapsedTimeScaled(), Valid: true},
		TotalTimerTime:   sql.NullFloat64{Float64: msg.GetTotalTimerTimeScaled(), Valid: true},
		TotalDistance:    sql.NullFloat64{Float64: msg.GetTotalDistanceScaled(), Valid: true},
		TotalCalories:    sql.NullInt16{Int16: int16(msg.TotalCalories), Valid: true},
		AvgSpeed:         sql.NullFloat64{Float64: msg.GetEnhancedAvgSpeedScaled(), Valid: true},
		MaxSpeed:         sql.NullFloat64{Float64: msg.GetEnhancedMaxSpeedScaled(), Valid: true},
		AvgHeartRate:     sql.NullInt16{Int16: int16(msg.AvgHeartRate), Valid: true},
		MaxHeartRate:     sql.NullInt16{Int16: int16(msg.MaxHeartRate), Valid: true},
		AvgVerticalRatio: sql.NullFloat64{Float64: msg.GetAvgVerticalRatioScaled(), Valid: true},
		AvgStanceTime:    sql.NullFloat64{Float64: msg.GetAvgStanceTimeScaled(), Valid: true},
	}
}

func createActivityRecordParams(activityID int64, msg *fit.RecordMsg) postgresql.CreateActivityRecordParams {
	return postgresql.CreateActivityRecordParams{
		Activity:            sql.NullInt64{Int64: activityID, Valid: true},
		Ts:                  sql.NullTime{Time: msg.Timestamp, Valid: true},
		Altitude:            sql.NullFloat64{Float64: msg.GetAltitudeScaled(), Valid: true},
		HeartRate:           sql.NullInt16{Int16: int16(msg.HeartRate), Valid: true},
		Cadence:             sql.NullInt16{Int16: int16(msg.Cadence), Valid: true},
		Distance:            sql.NullFloat64{Float64: msg.GetDistanceScaled(), Valid: true},
		Speed:               sql.NullFloat64{Float64: msg.GetSpeedScaled(), Valid: true},
		Cycles:              sql.NullInt16{Int16: int16(msg.Cycles), Valid: true},
		PositionLat:         sql.NullFloat64{Float64: msg.PositionLat.Degrees(), Valid: true},
		PositionLong:        sql.NullFloat64{Float64: msg.PositionLong.Degrees(), Valid: true},
		EnhancedAltitude:    sql.NullFloat64{Float64: msg.GetEnhancedAltitudeScaled(), Valid: true},
		EnhancedSpeed:       sql.NullFloat64{Float64: msg.GetEnhancedSpeedScaled(), Valid: true},
		LeftRightBalance:    sql.NullInt16{Int16: int16(msg.LeftRightBalance), Valid: true},
		GpsAccuracy:         sql.NullInt16{Int16: int16(msg.GpsAccuracy), Valid: true},
		VerticalOscillation: sql.NullFloat64{Float64: msg.GetVerticalOscillationScaled(), Valid: true},
		VerticalRatio:       sql.NullFloat64{Float64: msg.GetVerticalRatioScaled(), Valid: true},
		StanceTime:          sql.NullFloat64{Float64: msg.GetStanceTimeScaled(), Valid: true},
	}
}

func ingestActivity(ctx context.Context, queries *postgresql.Queries, data *fit.File, source string) error {
	activityFile, err := data.Activity()
	if err != nil {
		return err
	}

	params := createActivityParams(activityFile.Activity, source)
	log.Infof("Creating activity with params: %+v", params)

	id, err := queries.CreateActivity(ctx, params)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("failed to create activity: %w", err)
		}
	}

	if id == 0 { // No row inserted, skip inserting related tables.
		return nil
	}

	for i, session := range activityFile.Sessions {
		params := createActivitySessionParams(id, session)
		log.V(1).Infof("Creating session %d for activity %d with params: %+v", i, id, params)

		_, err := queries.CreateActivitySession(ctx, params)
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				return fmt.Errorf("failed to create session: %w", err)
			}
		}
	}

	for i, lap := range activityFile.Laps {
		params := createActivityLapParams(id, lap)
		log.V(1).Infof("Creating lap %d for activity %d with params: %+v", i, id, params)

		_, err := queries.CreateActivityLap(ctx, params)
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				return fmt.Errorf("failed to create lap: %w", err)
			}
		}
	}

	for i, record := range activityFile.Records {
		params := createActivityRecordParams(id, record)
		log.V(2).Infof("Creating record %d for activity %d with params: %+v", i, id, params)

		_, err := queries.CreateActivityRecord(ctx, params)
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				return fmt.Errorf("failed to create record: %w", err)
			}
		}
	}

	return nil
}

func ingestActivities(ctx context.Context, queries *postgresql.Queries) error {
	return filepath.WalkDir(filepath.Join(fitDir, "Activity"), func(path string, d fs.DirEntry, err error) error {
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

		return ingestActivity(ctx, queries, data, filepath.Base(path))
	})
}
