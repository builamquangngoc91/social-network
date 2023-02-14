package log

import (
	log "github.com/sirupsen/logrus"
)

func New() *log.Logger {
	l := log.New()
	l.SetFormatter(&log.JSONFormatter{})
	return l
} 