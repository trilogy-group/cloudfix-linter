package logger

import (
	"io/ioutil"
	"os"

	"github.com/google/uuid"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/writer"
)

var Log *logrus.Logger

func NewLogger(filePath string) *logrus.Logger {
	if Log != nil {
		return Log
	}
	flag := false
	fileName := ""
	if filePath == "" {
		flag = true
		id := uuid.New().String()
		filePath = "debug/Track-" + id + ".json"
		fileName = "Track-" + id + ".json"
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
	Log.SetOutput(ioutil.Discard) // Send all logs to nowhere by default
	Log.SetReportCaller(true)
	Log.Hooks.Add(lfshook.NewHook(
		pathMap,
		&logrus.JSONFormatter{},
	))
	Log.AddHook(&writer.Hook{ // Send logs with level higher than warning to stderr
		Writer: os.Stdout,
		LogLevels: []logrus.Level{
			logrus.InfoLevel,
			logrus.ErrorLevel,
			logrus.WarnLevel,
			logrus.FatalLevel,
			logrus.PanicLevel,
		},
	})
	Log.Info("Started Execution")
	if flag == true {
		Log.WithFields(logrus.Fields{
			"filename": fileName,
		}).Info("New File created in debug folder")
	}
	return Log
}
