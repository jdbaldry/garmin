package main

import (
	"context"
	"database/sql"
	_ "embed"
	"flag"
	"os"

	log "github.com/golang/glog"
	_ "github.com/lib/pq"

	"jdb.sh/garmin/postgresql"
)

//go:embed schema.sql
var ddl string

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

	if err := queries.PopulateMetadata(ctx); err != nil {
		log.Fatalln(err)
	}

	if err := queries.PopulateDashboards(ctx); err != nil {
		log.Fatalln(err)
	}

	if err := queries.PopulateSleepActivityLevels(ctx); err != nil {
		log.Fatalln(err)
	}

	if file := flag.Arg(0); file != "" {
		if err := ingest(ctx, queries, file); err != nil {
			log.Fatalln(err)
		}

		os.Exit(0)
	}

	if err := ingestActivities(ctx, queries); err != nil {
		log.Fatalln(err)
	}

	if err := ingestStressLevels(ctx, queries); err != nil {
		log.Fatalln(err)
	}

	if err := ingestSleeps(ctx, queries); err != nil {
		log.Fatalln(err)
	}

	if err := queries.PopulateActivitySessionsMetadata(ctx); err != nil {
		log.Fatalln(err)
	}
}
