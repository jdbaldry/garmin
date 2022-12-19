package main

import (
	"context"
	"io/fs"
	"path/filepath"

	log "github.com/golang/glog"

	"jdb.sh/garmin/postgresql"
)

func ingestSleeps(ctx context.Context, queries *postgresql.Queries) error {
	return filepath.WalkDir(filepath.Join(fitDir, "Sleep"), func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Errorln(err.Error())

			return nil // No errors for now, even though we can't ingest these files yet.
		}

		if !isFitFile(path) {
			return nil
		}

		_, err = decode(path)
		if err != nil {
			log.Errorln(err.Error())

			return nil // No errors for now, even though we can't ingest these files yet.
		}

		return nil
	})
}
