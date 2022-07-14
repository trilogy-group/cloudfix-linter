package logger

import (
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func NewLogger() *logrus.Logger {
	if Log != nil {
		return Log
	}

	pathMap := lfshook.PathMap{
		logrus.InfoLevel:  "/logs/console.log",
		logrus.ErrorLevel: "/logs/console.log",
		logrus.TraceLevel: "/logs/console.log",
		logrus.DebugLevel: "/logs/console.log",
	}

	Log = logrus.New()
	Log.Hooks.Add(lfshook.NewHook(
		pathMap,
		&logrus.JSONFormatter{},
	))
	return Log
}
