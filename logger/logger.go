package logger

import (
	"log"
	"github.com/google/uuid"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func NewLogger(filePath string) *logrus.Logger {
	if Log != nil {
		return Log
	}
	flag:=false
	fileName:=""
	if filePath == "" {
		flag=true
		id := uuid.New().String()
		filePath = "debug/Track-" + id + ".log"
		fileName = "Track-" + id + ".log"
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
    Log.Info("Started Execution")
	if flag == true {
    	log.Printf("New File \"%s\" created in debug folder\n", fileName)
	}
	return Log
}
