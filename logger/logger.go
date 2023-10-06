package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logrus.SetReportCaller(false)
}

func setOutPutFile(level logrus.Level, logName string) {
	if _, err := os.Stat("./runtime.log"); os.IsNotExist(err) {
		err = os.MkdirAll("./runtime.log", 0777)
		if err != nil {
			panic(fmt.Errorf("error %s", err))
		}
	}

	timeStr := time.Now().Format("2006-01-02")
	fileName := path.Join("./runtime/log", logName+"_"+timeStr+".log")

	var err error
	os.Stderr, err = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		fmt.Println("open log file err", err)
	}

	logrus.SetOutput(os.Stderr)
	logrus.SetLevel(level)
	return
}

func Write(msg string, filename string) {

	setOutPutFile(logrus.InfoLevel, filename)
	logrus.Info(msg)

}

func Debug(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.DebugLevel, "debug")
	logrus.WithFields(fields).Debug(args)
}

func Info(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.DebugLevel, "info")
	logrus.WithFields(fields).Info(args)
}

func Warn(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.DebugLevel, "info")
	logrus.WithFields(fields).Warn(args)
}

func Error(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.DebugLevel, "info")
	logrus.WithFields(fields).Error(args)
}
