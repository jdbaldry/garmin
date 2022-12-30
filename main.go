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

	log.Infoln("Populating sports")
	if err := queries.PopulateSports(ctx); err != nil {
		log.Fatalln(err)
	}

	log.Infoln("Populating sub sports")
	if err := queries.PopulateSubSports(ctx); err != nil {
		log.Fatalln(err)
	}

	log.Infoln("Populating metadata")
	if err := queries.PopulateMetadata(ctx); err != nil {
		log.Fatalln(err)
	}

	log.Infoln("Populating dashboards")
	if err := queries.PopulateDashboards(ctx); err != nil {
		log.Fatalln(err)
	}

	log.Infoln("Populating sleep activity levels")
	if err := queries.PopulateSleepActivityLevels(ctx); err != nil {
		log.Fatalln(err)
	}

	if file := flag.Arg(0); file != "" {
		if err := ingest(ctx, queries, file); err != nil {
			log.Fatalln(err)
		}

		os.Exit(0)
	}

	log.Infoln("Ingesting activities")
	if err := ingestActivities(ctx, queries); err != nil {
		log.Fatalln(err)
	}

	log.Infoln("Ingesting heart rates")
	if err := ingestHeartRates(ctx, queries); err != nil {
		log.Fatalln(err)
	}

	log.Infoln("Ingesting stress levels")
	if err := ingestStressLevels(ctx, queries); err != nil {
		log.Fatalln(err)
	}

	log.Infoln("Ingesting sleeps")
	if err := ingestSleeps(ctx, queries); err != nil {
		log.Fatalln(err)
	}

	if err := queries.PopulateActivitySessionsMetadata(ctx); err != nil {
		log.Fatalln(err)
	}
}
