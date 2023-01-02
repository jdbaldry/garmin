// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: query.sql

package postgresql

import (
	"context"
	"database/sql"
)

const createActivity = `-- name: CreateActivity :one
INSERT INTO activities (
    start_ts,
    end_ts,
    total_timer_time,
    num_sessions,
    type,
    event,
    event_type,
    local_ts,
    event_group,
    source
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
ON CONFLICT DO NOTHING
RETURNING id
`

type CreateActivityParams struct {
	StartTs        sql.NullTime
	EndTs          sql.NullTime
	TotalTimerTime sql.NullFloat64
	NumSessions    sql.NullInt32
	Type           sql.NullInt32
	Event          sql.NullInt16
	EventType      sql.NullInt16
	LocalTs        sql.NullTime
	EventGroup     sql.NullInt16
	Source         sql.NullString
}

func (q *Queries) CreateActivity(ctx context.Context, arg CreateActivityParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, createActivity,
		arg.StartTs,
		arg.EndTs,
		arg.TotalTimerTime,
		arg.NumSessions,
		arg.Type,
		arg.Event,
		arg.EventType,
		arg.LocalTs,
		arg.EventGroup,
		arg.Source,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const createActivityLap = `-- name: CreateActivityLap :one
INSERT INTO activity_laps (
    activity,
    message_index,
    start_ts,
    end_ts,
    event,
    event_type,
    sport,
    sub_sport,
    total_elapsed_time,
    total_timer_time,
    total_distance,
    total_calories,
    avg_speed,
    max_speed,
    avg_heart_rate,
    max_heart_rate,
    avg_vertical_ratio,
    avg_stance_time
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18)
ON CONFLICT DO NOTHING
RETURNING id
`

type CreateActivityLapParams struct {
	Activity         sql.NullInt64
	MessageIndex     sql.NullInt16
	StartTs          sql.NullTime
	EndTs            sql.NullTime
	Event            sql.NullInt16
	EventType        sql.NullInt16
	Sport            sql.NullInt16
	SubSport         sql.NullInt16
	TotalElapsedTime sql.NullFloat64
	TotalTimerTime   sql.NullFloat64
	TotalDistance    sql.NullFloat64
	TotalCalories    sql.NullInt16
	AvgSpeed         sql.NullFloat64
	MaxSpeed         sql.NullFloat64
	AvgHeartRate     sql.NullInt16
	MaxHeartRate     sql.NullInt16
	AvgVerticalRatio sql.NullFloat64
	AvgStanceTime    sql.NullFloat64
}

func (q *Queries) CreateActivityLap(ctx context.Context, arg CreateActivityLapParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, createActivityLap,
		arg.Activity,
		arg.MessageIndex,
		arg.StartTs,
		arg.EndTs,
		arg.Event,
		arg.EventType,
		arg.Sport,
		arg.SubSport,
		arg.TotalElapsedTime,
		arg.TotalTimerTime,
		arg.TotalDistance,
		arg.TotalCalories,
		arg.AvgSpeed,
		arg.MaxSpeed,
		arg.AvgHeartRate,
		arg.MaxHeartRate,
		arg.AvgVerticalRatio,
		arg.AvgStanceTime,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const createActivityRecord = `-- name: CreateActivityRecord :one
INSERT INTO activity_records (
    activity,
    ts,
    altitude,
    heart_rate,
    cadence,
    distance,
    speed,
    cycles,
    position_lat,
    position_long,
    enhanced_altitude,
    enhanced_speed,
    left_right_balance,
    gps_accuracy,
    vertical_oscillation,
    vertical_ratio,
    stance_time
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)
ON CONFLICT DO NOTHING
RETURNING id
`

type CreateActivityRecordParams struct {
	Activity            sql.NullInt64
	Ts                  sql.NullTime
	Altitude            sql.NullFloat64
	HeartRate           sql.NullInt16
	Cadence             sql.NullInt16
	Distance            sql.NullFloat64
	Speed               sql.NullFloat64
	Cycles              sql.NullInt16
	PositionLat         sql.NullFloat64
	PositionLong        sql.NullFloat64
	EnhancedAltitude    sql.NullFloat64
	EnhancedSpeed       sql.NullFloat64
	LeftRightBalance    sql.NullInt16
	GpsAccuracy         sql.NullInt16
	VerticalOscillation sql.NullFloat64
	VerticalRatio       sql.NullFloat64
	StanceTime          sql.NullFloat64
}

func (q *Queries) CreateActivityRecord(ctx context.Context, arg CreateActivityRecordParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, createActivityRecord,
		arg.Activity,
		arg.Ts,
		arg.Altitude,
		arg.HeartRate,
		arg.Cadence,
		arg.Distance,
		arg.Speed,
		arg.Cycles,
		arg.PositionLat,
		arg.PositionLong,
		arg.EnhancedAltitude,
		arg.EnhancedSpeed,
		arg.LeftRightBalance,
		arg.GpsAccuracy,
		arg.VerticalOscillation,
		arg.VerticalRatio,
		arg.StanceTime,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const createActivitySession = `-- name: CreateActivitySession :one
INSERT INTO activity_sessions (
    activity,
    start_ts,
    end_ts,
    event,
    event_type,
    sport,
    sub_sport,
    total_elapsed_time,
    total_timer_time,
    total_distance,
    total_calories,
    avg_speed,
    max_speed,
    avg_heart_rate,
    max_heart_rate,
    avg_vertical_ratio,
    avg_stance_time
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)
ON CONFLICT DO NOTHING
RETURNING id
`

type CreateActivitySessionParams struct {
	Activity         sql.NullInt64
	StartTs          sql.NullTime
	EndTs            sql.NullTime
	Event            sql.NullInt16
	EventType        sql.NullInt16
	Sport            sql.NullInt16
	SubSport         sql.NullInt16
	TotalElapsedTime sql.NullFloat64
	TotalTimerTime   sql.NullFloat64
	TotalDistance    sql.NullFloat64
	TotalCalories    sql.NullInt16
	AvgSpeed         sql.NullFloat64
	MaxSpeed         sql.NullFloat64
	AvgHeartRate     sql.NullInt16
	MaxHeartRate     sql.NullInt16
	AvgVerticalRatio sql.NullFloat64
	AvgStanceTime    sql.NullFloat64
}

func (q *Queries) CreateActivitySession(ctx context.Context, arg CreateActivitySessionParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, createActivitySession,
		arg.Activity,
		arg.StartTs,
		arg.EndTs,
		arg.Event,
		arg.EventType,
		arg.Sport,
		arg.SubSport,
		arg.TotalElapsedTime,
		arg.TotalTimerTime,
		arg.TotalDistance,
		arg.TotalCalories,
		arg.AvgSpeed,
		arg.MaxSpeed,
		arg.AvgHeartRate,
		arg.MaxHeartRate,
		arg.AvgVerticalRatio,
		arg.AvgStanceTime,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const createHeartRate = `-- name: CreateHeartRate :one
INSERT INTO heart_rates (ts, value)
VALUES ($1, $2)
ON CONFLICT DO NOTHING
RETURNING id
`

type CreateHeartRateParams struct {
	Ts    sql.NullTime
	Value sql.NullInt16
}

func (q *Queries) CreateHeartRate(ctx context.Context, arg CreateHeartRateParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, createHeartRate, arg.Ts, arg.Value)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const createMonitoring = `-- name: CreateMonitoring :one
INSERT INTO monitorings (
  ts,
  calories,
  cycles,
  distance,
  active_time,
  activity_type,
  activity_sub_type,
  local_ts
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
ON CONFLICT DO NOTHING
RETURNING id
`

type CreateMonitoringParams struct {
	Ts              sql.NullTime
	Calories        sql.NullInt16
	Cycles          sql.NullInt32
	Distance        sql.NullFloat64
	ActiveTime      sql.NullFloat64
	ActivityType    sql.NullInt16
	ActivitySubType sql.NullInt16
	LocalTs         sql.NullTime
}

func (q *Queries) CreateMonitoring(ctx context.Context, arg CreateMonitoringParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, createMonitoring,
		arg.Ts,
		arg.Calories,
		arg.Cycles,
		arg.Distance,
		arg.ActiveTime,
		arg.ActivityType,
		arg.ActivitySubType,
		arg.LocalTs,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const createRecord = `-- name: CreateRecord :one
INSERT INTO records (distance, time)
VALUES ($1, $2)
ON CONFLICT DO NOTHING
RETURNING id
`

type CreateRecordParams struct {
	Distance sql.NullInt32
	Time     sql.NullInt32
}

func (q *Queries) CreateRecord(ctx context.Context, arg CreateRecordParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, createRecord, arg.Distance, arg.Time)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const createSleep = `-- name: CreateSleep :one
INSERT INTO sleeps (
  start_ts,
  end_ts
)
VALUES ($1, $2)
ON CONFLICT DO NOTHING
RETURNING id
`

type CreateSleepParams struct {
	StartTs sql.NullTime
	EndTs   sql.NullTime
}

func (q *Queries) CreateSleep(ctx context.Context, arg CreateSleepParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, createSleep, arg.StartTs, arg.EndTs)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const createSleepRecord = `-- name: CreateSleepRecord :one
INSERT INTO sleep_records (
  sleep,
  start_ts,
  end_ts,
  sleep_activity_level
)
VALUES ($1, $2, $3, $4)
ON CONFLICT DO NOTHING
RETURNING id
`

type CreateSleepRecordParams struct {
	Sleep              sql.NullInt64
	StartTs            sql.NullTime
	EndTs              sql.NullTime
	SleepActivityLevel sql.NullInt16
}

func (q *Queries) CreateSleepRecord(ctx context.Context, arg CreateSleepRecordParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, createSleepRecord,
		arg.Sleep,
		arg.StartTs,
		arg.EndTs,
		arg.SleepActivityLevel,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const createStressLevel = `-- name: CreateStressLevel :one
INSERT INTO stress_levels (ts, value)
VALUES ($1, $2)
ON CONFLICT DO NOTHING
RETURNING id
`

type CreateStressLevelParams struct {
	Ts    sql.NullTime
	Value sql.NullInt16
}

func (q *Queries) CreateStressLevel(ctx context.Context, arg CreateStressLevelParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, createStressLevel, arg.Ts, arg.Value)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const getActivity = `-- name: GetActivity :one
SELECT id, start_ts, end_ts, total_timer_time, num_sessions, type, event, event_type, local_ts, event_group, source FROM activities
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetActivity(ctx context.Context, id int64) (Activity, error) {
	row := q.db.QueryRowContext(ctx, getActivity, id)
	var i Activity
	err := row.Scan(
		&i.ID,
		&i.StartTs,
		&i.EndTs,
		&i.TotalTimerTime,
		&i.NumSessions,
		&i.Type,
		&i.Event,
		&i.EventType,
		&i.LocalTs,
		&i.EventGroup,
		&i.Source,
	)
	return i, err
}

const populateActivitySessionsMetadata = `-- name: PopulateActivitySessionsMetadata :exec
INSERT INTO activity_sessions_metadata (activity_session, kind, value)
VALUES
  (216, 1, 'Thorpe Marriot'),
  (216, 2, 'Nearly a half marathon with Mabel'),
  (209, 1, 'Sheringham Park with Alisha'),
  (208, 1, 'Martham base'),
  (205, 1, 'From Highball'),
  (203, 1, 'To Highball'),
  (201, 1, 'Wednesday night football'),
  (201, 2, 'https://football.jdb.sh/2022/2022-12-21.html'),
  (200, 1, 'Testing Football.PRG'),
  (127, 1, 'Grinnell glacier')
ON CONFLICT DO NOTHING
`

func (q *Queries) PopulateActivitySessionsMetadata(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, populateActivitySessionsMetadata)
	return err
}

const populateDashboards = `-- name: PopulateDashboards :exec
INSERT INTO dashboards (sport, uid, title)
VALUES
  (1, 'nzZ73htVz', 'running'),
  (11, '9wmcqhpVk', 'walking'),
  (17, 'gotNq2pVz', 'hiking'),
  (31, 'MfI_jhp4k', 'rock-climbing'),
  (41, 'Y0hvq2p4z', 'kayaking')
ON CONFLICT DO NOTHING
`

func (q *Queries) PopulateDashboards(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, populateDashboards)
	return err
}

const populateMetadata = `-- name: PopulateMetadata :exec
INSERT INTO metadata (id, name)
VALUES
  (1, 'name'),
  (2, 'comment'),
  (3, 'source')
ON CONFLICT DO NOTHING
`

func (q *Queries) PopulateMetadata(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, populateMetadata)
	return err
}

const populateSleepActivityLevels = `-- name: PopulateSleepActivityLevels :exec
INSERT INTO sleep_activity_levels (id, name)
VALUES
  (0, 'Unmeasurable'),
  (1, 'Awake'),
  (2, 'Light'),
  (3, 'Deep'),
  (4, 'REM')
ON CONFLICT DO NOTHING
`

func (q *Queries) PopulateSleepActivityLevels(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, populateSleepActivityLevels)
	return err
}

const populateSports = `-- name: PopulateSports :exec
INSERT INTO sports (id, name)
VALUES
(0, 'Generic'),
(1, 'Running'),
(2, 'Cycling'),
(3, 'Transition'),
(4, 'Fitness equipment'),
(5, 'Swimming'),
(6, 'Basketball'),
(7, 'Soccer'),
(8, 'Tennis'),
(9, 'American football'),
(10, 'Training'),
(11, 'Walking'),
(12, 'Cross country skiing'),
(13, 'Alpine skiing'),
(14, 'Snowboarding'),
(15, 'Rowing'),
(16, 'Mountaineering'),
(17, 'Hiking'),
(18, 'Multisport'),
(19, 'Paddling'),
(20, 'Flying'),
(21, 'EBiking'),
(22, 'Motorcycling'),
(23, 'Boating'),
(24, 'Driving'),
(25, 'Golf'),
(26, 'Hang gliding'),
(27, 'Horseback riding'),
(28, 'Hunting'),
(29, 'Fishing'),
(30, 'Inline skating'),
(31, 'Rock climbing'),
(32, 'Sailing'),
(33, 'Ice skating'),
(34, 'Sky diving'),
(35, 'Snowshoeing'),
(36, 'Snowmobiling'),
(37, 'Stand up paddleboarding'),
(38, 'Surfing'),
(39, 'Wakeboarding'),
(40, 'Water skiing'),
(41, 'Kayaking'),
(42, 'Rafting'),
(43, 'Windsurfing'),
(44, 'Kitesurfing'),
(45, 'Tactical'),
(46, 'Jumpmaster'),
(47, 'Sport boxing'),
(48, 'Floor climbing'),
(49, 'Sleep'),
(53, 'Diving'),
(254, 'All'),
(255, 'Invalid')
ON CONFLICT DO NOTHING
`

func (q *Queries) PopulateSports(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, populateSports)
	return err
}

const populateSubSports = `-- name: PopulateSubSports :exec
INSERT INTO sub_sports (id, name)
VALUES
  (0, 'Generic'),
  (1, 'Treadmill'),  -- Run/Fitness Equipment
  (2, 'Street'),  -- Run
  (3, 'Trail'),  -- Run
  (4, 'Track'),  -- Run
  (5, 'Spin'),  -- Cycling
  (6, 'Indoor cycling'),  -- Cycling/Fitness Equipment
  (7, 'Road'),  -- Cycling
  (8, 'Mountain'),  -- Cycling
  (9, 'Downhill'),  -- Cycling
  (10, 'Recumbent'), -- Cycling
  (11, 'Cyclocross'), -- Cycling
  (12, 'Hand cycling'), -- Cycling
  (13, 'Track cycling'), -- Cycling
  (14, 'Indoor rowing'), -- Fitness Equipment
  (15, 'Elliptical'), -- Fitness Equipment
  (16, 'Stair climbing'), -- Fitness Equipment
  (17, 'Lap swimming'), -- Swimming
  (18, 'Open water'), -- Swimming
  (19, 'Flexibility training'), -- Training
  (20, 'Strength training'), -- Training
  (21, 'Warm up'), -- Tennis
  (22, 'Match'), -- Tennis
  (23, 'Exercise'), -- Tennis
  (24, 'Challenge'),
  (25, 'Indoor skiing'), -- Fitness Equipment
  (26, 'Cardio training'), -- Training
  (27, 'Indoor walking'), -- Walking/Fitness Equipment
  (28, 'EBike fitness'), -- E-Biking
  (29, 'BMX'), -- Cycling
  (30, 'Casual walking'), -- Walking
  (31, 'Speed walking'), -- Walking
  (32, 'Bike to run transition'), -- Transition
  (33, 'Run to bike transition'), -- Transition
  (34, 'Swim to bike transition'), -- Transition
  (35, 'ATV'), -- Motorcycling
  (36, 'Motocross'), -- Motorcycling
  (37, 'Backcountry'), -- Alpine Skiing/Snowboarding
  (38, 'Resort'), -- Alpine Skiing/Snowboarding
  (39, 'RC drone'), -- Flying
  (40, 'Wingsuit'), -- Flying
  (41, 'Whitewater'), -- Kayaking/Rafting
  (42, 'Skate skiing'), -- Cross Country Skiing
  (43, 'Yoga'), -- Training
  (44, 'Pilates'), -- Fitness Equipment
  (45, 'Indoor running'), -- Run
  (46, 'Gravel cycling'), -- Cycling
  (47, 'EBike mountain'), -- Cycling
  (48, 'Commuting'), -- Cycling
  (49, 'Mixed surface'), -- Cycling
  (50, 'Navigate'),
  (51, 'Track me'),
  (52, 'Map'),
  (53, 'Single gas diving'), -- Diving
  (54, 'Multi gas diving'), -- Diving
  (55, 'Gauge diving'), -- Diving
  (56, 'Apnea diving'), -- Diving
  (57, 'Apnea hunting'), -- Diving
  (58, 'Virtual activity'),
  (59, 'Obstacle'), -- Used for events where participants run, crawl through mud, climb over walls, etc.
  (62, 'Breathing'),
  (65, 'Sail race'), -- Sailing
  (67, 'Ultra'), -- Ultramarathon
  (68, 'Indoor climbing'), -- Climbing
  (69, 'Bouldering'), -- Climbing
  (254, 'All'),
  (255, 'Invalid')
ON CONFLICT DO NOTHING
`

func (q *Queries) PopulateSubSports(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, populateSubSports)
	return err
}
