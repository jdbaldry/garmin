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

var errNotActivity = errors.New("not an activity")

func createActivityParams(activity *fit.ActivityMsg, source string) postgresql.CreateActivityParams {
	return postgresql.CreateActivityParams{
		StartTs: sql.NullTime{
			Time:  activity.Timestamp.Add(-(time.Duration(activity.GetTotalTimerTimeScaled()) * time.Second)),
			Valid: true,
		},
		EndTs:          sql.NullTime{Time: activity.Timestamp, Valid: true},
		TotalTimerTime: sql.NullFloat64{Float64: activity.GetTotalTimerTimeScaled(), Valid: true},
		NumSessions:    sql.NullInt32{Int32: int32(activity.NumSessions), Valid: true},
		Type:           sql.NullInt32{Int32: int32(activity.Type), Valid: true},
		Event:          sql.NullInt16{Int16: int16(activity.Event), Valid: true},
		EventType:      sql.NullInt16{Int16: int16(activity.EventType), Valid: true},
		LocalTs:        sql.NullTime{Time: activity.LocalTimestamp, Valid: true},
		EventGroup:     sql.NullInt16{Int16: int16(activity.EventGroup), Valid: true},
		Source:         sql.NullString{String: source, Valid: true},
	}
}

func createActivitySessionParams(activityID int64, session *fit.SessionMsg) postgresql.CreateActivitySessionParams {
	return postgresql.CreateActivitySessionParams{
		Activity:         sql.NullInt64{Int64: activityID, Valid: true},
		StartTs:          sql.NullTime{Time: session.StartTime, Valid: true},
		EndTs:            sql.NullTime{Time: session.Timestamp, Valid: true},
		Event:            sql.NullInt16{Int16: int16(session.Event), Valid: true},
		EventType:        sql.NullInt16{Int16: int16(session.EventType), Valid: true},
		Sport:            sql.NullInt16{Int16: int16(session.Sport), Valid: true},
		SubSport:         sql.NullInt16{Int16: int16(session.SubSport), Valid: true},
		TotalElapsedTime: sql.NullFloat64{Float64: session.GetTotalElapsedTimeScaled(), Valid: true},
		TotalTimerTime:   sql.NullFloat64{Float64: session.GetTotalTimerTimeScaled(), Valid: true},
		TotalDistance:    sql.NullFloat64{Float64: session.GetTotalDistanceScaled(), Valid: true},
		TotalCalories:    sql.NullInt16{Int16: int16(session.TotalCalories), Valid: true},
		AvgSpeed:         sql.NullFloat64{Float64: session.GetEnhancedAvgSpeedScaled(), Valid: true},
		MaxSpeed:         sql.NullFloat64{Float64: session.GetEnhancedMaxSpeedScaled(), Valid: true},
		AvgHeartRate:     sql.NullInt16{Int16: int16(session.AvgHeartRate), Valid: true},
		MaxHeartRate:     sql.NullInt16{Int16: int16(session.MaxHeartRate), Valid: true},
	}
}

func createActivityLapParams(activityID int64, lap *fit.LapMsg) postgresql.CreateActivityLapParams {
	return postgresql.CreateActivityLapParams{
		Activity:         sql.NullInt64{Int64: activityID, Valid: true},
		MessageIndex:     sql.NullInt16{Int16: 0, Valid: false},
		StartTs:          sql.NullTime{Time: lap.StartTime, Valid: true},
		EndTs:            sql.NullTime{Time: lap.Timestamp, Valid: true},
		Event:            sql.NullInt16{Int16: int16(lap.Event), Valid: true},
		EventType:        sql.NullInt16{Int16: int16(lap.EventType), Valid: true},
		Sport:            sql.NullInt16{Int16: int16(lap.Sport), Valid: true},
		SubSport:         sql.NullInt16{Int16: int16(lap.SubSport), Valid: true},
		TotalElapsedTime: sql.NullFloat64{Float64: lap.GetTotalElapsedTimeScaled(), Valid: true},
		TotalTimerTime:   sql.NullFloat64{Float64: lap.GetTotalTimerTimeScaled(), Valid: true},
		TotalDistance:    sql.NullFloat64{Float64: lap.GetTotalDistanceScaled(), Valid: true},
		TotalCalories:    sql.NullInt16{Int16: int16(lap.TotalCalories), Valid: true},
		AvgSpeed:         sql.NullFloat64{Float64: lap.GetEnhancedAvgSpeedScaled(), Valid: true},
		MaxSpeed:         sql.NullFloat64{Float64: lap.GetEnhancedMaxSpeedScaled(), Valid: true},
		AvgHeartRate:     sql.NullInt16{Int16: int16(lap.AvgHeartRate), Valid: true},
		MaxHeartRate:     sql.NullInt16{Int16: int16(lap.MaxHeartRate), Valid: true},
	}
}

func createActivityRecordParams(activityID int64, record *fit.RecordMsg) postgresql.CreateActivityRecordParams {
	return postgresql.CreateActivityRecordParams{
		Activity:            sql.NullInt64{Int64: activityID, Valid: true},
		Ts:                  sql.NullTime{Time: record.Timestamp, Valid: true},
		Altitude:            sql.NullFloat64{Float64: record.GetAltitudeScaled(), Valid: true},
		HeartRate:           sql.NullInt16{Int16: int16(record.HeartRate), Valid: true},
		Cadence:             sql.NullInt16{Int16: int16(record.Cadence), Valid: true},
		Distance:            sql.NullFloat64{Float64: record.GetDistanceScaled(), Valid: true},
		Speed:               sql.NullFloat64{Float64: record.GetSpeedScaled(), Valid: true},
		Cycles:              sql.NullInt16{Int16: int16(record.Cycles), Valid: true},
		PositionLat:         sql.NullFloat64{Float64: record.PositionLat.Degrees(), Valid: true},
		PositionLong:        sql.NullFloat64{Float64: record.PositionLong.Degrees(), Valid: true},
		EnhancedAltitude:    sql.NullFloat64{Float64: record.GetEnhancedAltitudeScaled(), Valid: true},
		EnhancedSpeed:       sql.NullFloat64{Float64: record.GetEnhancedSpeedScaled(), Valid: true},
		LeftRightBalance:    sql.NullInt16{Int16: int16(record.LeftRightBalance), Valid: true},
		GpsAccuracy:         sql.NullInt16{Int16: int16(record.GpsAccuracy), Valid: true},
		VerticalOscillation: sql.NullFloat64{Float64: record.GetVerticalOscillationScaled(), Valid: true},
	}
}

func ingestActivity(ctx context.Context, queries *postgresql.Queries, data *fit.File, source string) error {
	activityFile, err := data.Activity()
	if err != nil {
		return errNotActivity
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
