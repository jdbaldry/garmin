package main

import (
	"context"
	"fmt"
	"path/filepath"

	"jdb.sh/garmin/postgresql"
)

func ingest(ctx context.Context, queries *postgresql.Queries, path string) error {
	if !isFitFile(path) {
		return fmt.Errorf("unable to ingest %q: %w", path, errNotFitFile)
	}

	data, err := decode(path)
	if err != nil {
		return err
	}

	return ingestActivity(ctx, queries, data, filepath.Base(path))
}
