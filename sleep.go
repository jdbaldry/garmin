package main

import (
	"context"
	"io/fs"
	"path/filepath"

	"jdb.sh/garmin/postgresql"
)

func ingestSleeps(ctx context.Context, queries *postgresql.Queries) error {
	return filepath.WalkDir(filepath.Join(fitDir, "Sleep"), func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !isFitFile(path) {
			return nil
		}

		_, err = decode(path)
		if err != nil {
			return err
		}

		return nil
	})
}
