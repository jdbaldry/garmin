package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/tormoder/fit"

	"jdb.sh/garmin/postgresql"
)

const fitDir = "fit/Primary/GARMIN"

var errNotFitFile = errors.New("not a FIT file")

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
