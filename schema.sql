-- [[file:vendor/github.com/tormoder/fit/messages.go::// ActivityMsg represents the activity FIT message type.][Activities]]
CREATE TABLE IF NOT EXISTS activities (
    id bigserial PRIMARY KEY,
    start_ts timestamp,
    end_ts timestamp,
    total_timer_time double precision, -- double precision is used because the scaled values returns a float64
    num_sessions integer,
    type integer,
    event smallint,
    event_type smallint,
    local_ts timestamp,
    event_group smallint,
    source varchar,
    UNIQUE (event, start_ts, end_ts)
);

-- [[file:vendor/github.com/tormoder/fit/messages.go::// SessionMsg represents the session FIT message type.][Sessions]]
CREATE TABLE IF NOT EXISTS activity_sessions (
    id bigserial PRIMARY KEY,
    activity bigint REFERENCES activities(id),
    start_ts timestamp,
    end_ts timestamp,
    event smallint,
    event_type smallint,
    sport smallint,
    sub_sport smallint,
    total_elapsed_time double precision, -- double precision is used because the scaled values returns a float64
    total_timer_time double precision, -- double precision is used because the scaled values returns a float64
    total_distance double precision, -- double precision is used because the scaled values returns a float64
    total_calories smallint,
    avg_speed double precision, -- double precision is used because the scaled values returns a float64
    max_speed double precision, -- double precision is used because the scaled values returns a float64
    avg_heart_rate smallint,
    max_heart_rate smallint,
    avg_vertical_ratio double precision, -- double precision is used because the scaled values returns a float64
    avg_stance_time double precision, -- double precision is used because the scaled values returns a float64
    UNIQUE (activity, start_ts, end_ts)
);

-- [[file:vendor/github.com/tormoder/fit/messages.go::// LapMsg represents the lap FIT message type.][Laps]]
CREATE TABLE IF NOT EXISTS activity_laps (
    id bigserial PRIMARY KEY,
    activity bigint REFERENCES activities(id),
    message_index smallint,
    start_ts timestamp,
    end_ts timestamp,
    event smallint,
    event_type smallint,
    sport smallint,
    sub_sport smallint,
    total_elapsed_time double precision, -- double precision is used because the scaled values returns a float64
    total_timer_time double precision, -- double precision is used because the scaled values returns a float64
    total_distance double precision, -- double precision is used because the scaled values returns a float64
    total_calories smallint,
    avg_speed double precision, -- double precision is used because the scaled values returns a float64
    max_speed double precision, -- double precision is used because the scaled values returns a float64
    avg_heart_rate smallint,
    max_heart_rate smallint,
    avg_vertical_ratio double precision, -- double precision is used because the scaled values returns a float64
    avg_stance_time double precision, -- double precision is used because the scaled values returns a float64
    UNIQUE (activity, start_ts, end_ts)
);

-- [[file:vendor/github.com/tormoder/fit/messages.go::// RecordMsg represents the record FIT message type.][Records]]
CREATE TABLE IF NOT EXISTS activity_records (
    id bigserial PRIMARY KEY,
    activity bigint REFERENCES activities(id),
    ts timestamp,
    altitude double precision,  -- double precision is used because the scaled values returns a float64
    heart_rate smallint,
    cadence smallint,
    distance double precision, -- double precision is used because the scaled values returns a float64
    speed double precision, -- double precision is used because the scaled values returns a float64
    cycles smallint,
    position_lat double precision,  -- double precision is used because the fit library uses a float64 representation for degrees
    position_long double precision, -- double precision is used because the fit library uses a float64 representation for degrees
    enhanced_altitude double precision, -- double precision is used because the scaled values returns a float64
    enhanced_speed double precision, -- double precision is used because the scaled values returns a float64
    left_right_balance smallint,
    gps_accuracy smallint,
    vertical_oscillation double precision, -- double precision is used because the scaled values returns a float64
    vertical_ratio double precision, -- double precision is used because the scaled values returns a float64
    stance_time double precision, -- double precision is used because the scaled values returns a float64
    UNIQUE (activity, ts)
);

-- [[file:vendor/github.com/tormoder/fit/messages.go::// MonitoringMsg represents the monitoring FIT message type.][Monitorings]]
CREATE TABLE IF NOT EXISTS monitorings (
  id bigserial PRIMARY KEY,
  ts timestamp,
  cycles integer,
  calories smallint,
  distance double precision, -- double precision is used because the scaled values return a float64
  active_time double precision,  -- double precision is used because the scaled values return a float64
  activity_type smallint,
  activity_sub_type smallint,
  local_ts timestamp,
  UNIQUE (ts, activity_type, activity_sub_type)
);

-- [[file:csv/RECORDS/RECORDS.md::Records][Records]]
CREATE TABLE IF NOT EXISTS records (
  id bigserial PRIMARY KEY,
  distance integer,
  time integer
);

CREATE TABLE IF NOT EXISTS sports (
  id smallint PRIMARY KEY,
  name varchar
);

CREATE TABLE IF NOT EXISTS sub_sports (
  id smallint PRIMARY KEY,
  name varchar
);

CREATE TABLE IF NOT EXISTS sleeps (
  id bigserial PRIMARY KEY,
  start_ts timestamp,
  end_ts timestamp,
  UNIQUE (start_ts, end_ts)
);

CREATE TABLE IF NOT EXISTS sleep_records (
  id bigserial PRIMARY KEY,
  sleep bigint REFERENCES sleeps(id),
  start_ts timestamp,
  end_ts timestamp,
  sleep_activity_level smallint REFERENCES sleep_activity_levels(id)
);

CREATE TABLE IF NOT EXISTS sleep_activity_levels (
  id smallint PRIMARY KEY,
  name varchar
);

CREATE TABLE IF NOT EXISTS metadata (
  id bigserial PRIMARY KEY,
  name varchar
);

CREATE TABLE IF NOT EXISTS activity_sessions_metadata (
  id bigserial PRIMARY KEY,
  activity_session bigint REFERENCES activity_sessions(id),
  kind bigint REFERENCES metadata(id),
  value varchar,
  UNIQUE (activity_session, kind, value)
);

CREATE TABLE IF NOT EXISTS dashboards (
  sport bigint REFERENCES sports(id),
  uid varchar,
  title varchar
);

CREATE TABLE IF NOT EXISTS stress_levels (
  id bigserial PRIMARY KEY,
  ts timestamp UNIQUE,
  value smallint
);

CREATE TABLE IF NOT EXISTS heart_rates (
  id bigserial PRIMARY KEY,
  ts timestamp UNIQUE,
  value smallint
);

