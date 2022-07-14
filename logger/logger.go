package logger

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func NewLogger() *logrus.Logger {
	if Log != nil {
		return Log
	}
	var flag string
	var filePath string
	fmt.Println("Enter Path to log debug files: (Y/N)")
	fmt.Scan(&flag)
	if flag == "N" {
		id := uuid.New().String()
		filePath = "debug/Track-" + id + ".log"
		fileName := "Track-" + id + ".log"
		fmt.Printf("New File \"%s\" created in debug folder\n", fileName)
	} else {
		fmt.Scan(&filePath)
	}
	pathMap := lfshook.PathMap{
		logrus.InfoLevel:  filePath,
		logrus.ErrorLevel: filePath,
		logrus.TraceLevel: filePath,
		logrus.DebugLevel: filePath,
		logrus.PanicLevel: filePath,
		logrus.WarnLevel:  filePath,
		logrus.FatalLevel: filePath,
	}
	Log = logrus.New()
	Log.SetReportCaller(true)
	Log.SetLevel(logrus.TraceLevel)
	Log.Hooks.Add(lfshook.NewHook(
		pathMap,
		&logrus.JSONFormatter{},
	))
	return Log
}
