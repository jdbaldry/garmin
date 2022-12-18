package main

import (
	"context"
	"database/sql"
	_ "embed"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	log "github.com/golang/glog"
	_ "github.com/lib/pq"
	"github.com/tormoder/fit"

	"jdb.sh/garmin/postgresql"
)

const fitDir = "fit/Primary/GARMIN"

//go:embed schema.sql
var ddl string

func decode(path string) (*fit.File, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open file %s: %w", path, err)
	}

	data, err := fit.Decode(f, fit.WithUnknownMessages())
	if err != nil {
		return nil, fmt.Errorf("unable to decode FIT data in file %s: %w", path, err)
	}

	return data, nil
}

func isFitFile(path string) bool {
	suffix := filepath.Ext(path)
	if suffix != ".fit" && suffix != ".FIT" {
		return false
	}

	return true
}

func main() {
	flag.Parse()

	ctx := context.Background()

	db, err := sql.Open("postgres", "dbname=garmin password=garmin sslmode=disable user=garmin")
	if err != nil {
		log.Fatalln(err)
	}

	if _, err := db.ExecContext(ctx, ddl); err != nil {
		log.Fatalln(err)
	}

	queries := postgresql.New(db)

	if err := queries.PopulateSports(ctx); err != nil {
		log.Fatalln(err)
	}

	if err := queries.PopulateSubSports(ctx); err != nil {
		log.Fatalln(err)
	}

	if err := ingestActivities(ctx, queries); err != nil {
		log.Fatalln(err)
	}

	if err := ingestSleeps(ctx, queries); err != nil {
		log.Fatalln(err)
	}
}
