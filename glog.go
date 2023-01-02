package main

import (
	log "github.com/golang/glog"
)

type glog struct{}

// Print implements fit.Logger.
func (g *glog) Print(args ...interface{}) {
	log.Info(args...)
}

// Printf implements fit.Logger.
func (g *glog) Printf(format string, args ...interface{}) {
	log.Infof(format, args...)
}

// Println implements fit.Logger.
func (g *glog) Println(args ...interface{}) {
	log.Infoln(args...)
}
