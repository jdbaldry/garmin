// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package postgresql

import (
	"database/sql"
)

type Activity struct {
	ID             int64
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

type ActivityLap struct {
	ID               int64
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

type ActivityRecord struct {
	ID                  int64
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

type ActivitySession struct {
	ID               int64
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

type ActivitySessionsMetadatum struct {
	ID              int64
	ActivitySession sql.NullInt64
	Kind            sql.NullInt64
	Value           sql.NullString
}

type Dashboard struct {
	Sport sql.NullInt64
	Uid   sql.NullString
	Title sql.NullString
}

type HeartRate struct {
	ID    int64
	Ts    sql.NullTime
	Value sql.NullInt16
}

type Metadatum struct {
	ID   int64
	Name sql.NullString
}

type Monitoring struct {
	ID              int64
	Ts              sql.NullTime
	Cycles          sql.NullInt32
	Calories        sql.NullInt16
	Distance        sql.NullFloat64
	ActiveTime      sql.NullFloat64
	ActivityType    sql.NullInt16
	ActivitySubType sql.NullInt16
	LocalTs         sql.NullTime
}

type Record struct {
	ID       int64
	Distance sql.NullInt32
	Time     sql.NullInt32
}

type Sleep struct {
	ID      int64
	StartTs sql.NullTime
	EndTs   sql.NullTime
}

type SleepActivityLevel struct {
	ID   int16
	Name sql.NullString
}

type SleepRecord struct {
	ID                 int64
	Sleep              sql.NullInt64
	StartTs            sql.NullTime
	EndTs              sql.NullTime
	SleepActivityLevel sql.NullInt16
}

type Sport struct {
	ID   int16
	Name sql.NullString
}

type Step struct {
	ID              int64
	Ts              sql.NullTime
	Distance        sql.NullFloat64
	Cycles          sql.NullFloat64
	ActiveTime      sql.NullFloat64
	ActiveCalories  sql.NullInt32
	DurationMin     sql.NullInt16
	ActivityType    sql.NullInt16
	ActivitySubType sql.NullInt16
}

type StepsDaily struct {
	Day          int32
	Cycles       interface{}
	Distance     interface{}
	ActivityType sql.NullInt16
}

type StressLevel struct {
	ID    int64
	Ts    sql.NullTime
	Value sql.NullInt16
}

type SubSport struct {
	ID   int16
	Name sql.NullString
}
