package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/tormoder/fit"
)

const (
	// Difference between UNIX epoch and Garmin epoch in seconds.
	epochOffset = 631_065_600
)

const fitDir = "fit/Primary/GARMIN"

var errNotFitFile = errors.New("not a FIT file")

func decode(path string) (*fit.File, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open file %s: %w", path, err)
	}

	data, err := fit.Decode(f, fit.WithUnknownMessages(), fit.WithUnknownFields(), fit.WithLogger(&glog{}))
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
