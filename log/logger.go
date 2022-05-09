package log

import (
	"fmt"
	_ "gin_demo/config"
	"log"
	"os"
	"path"

	"github.com/spf13/viper"
)

type Logger struct {
	*log.Logger
}

var Log Logger

func init() {
	logPath := viper.GetString("LOG_PATH")
	logFileName := viper.GetString("LOG_FILENAME")
	logFile := path.Join(logPath, logFileName)
	fd, err := os.Create(logFile)
	if err != nil {
		panic(fmt.Errorf("Fatal error init log file: %s \n", err))
	}
	Log = Logger{log.New(fd, "", log.LstdFlags|log.Lshortfile)}
}

func (l Logger) Debug(v ...interface{}) {
	l.SetPrefix("debug: ")
	l.Print(v...)
}
func (l Logger) Debugf(format string, v ...interface{}) {
	l.SetPrefix("debug: ")
	l.Printf(format, v...)
}
func (l Logger) Info(v ...interface{}) {
	l.SetPrefix("info: ")
	l.Print(v...)
}
func (l Logger) Infof(format string, v ...interface{}) {
	l.SetPrefix("info: ")
	l.Printf(format, v...)
}
func (l Logger) Warn(v ...interface{}) {
	l.SetPrefix("warn: ")
	l.Print(v...)
}
func (l Logger) Warnf(format string, v ...interface{}) {
	l.SetPrefix("warn: ")
	l.Printf(format, v...)
}
func (l Logger) Error(v ...interface{}) {
	l.SetPrefix("error: ")
	l.Print(v...)
}
func (l Logger) Errorf(format string, v ...interface{}) {
	l.SetPrefix("error: ")
	l.Printf(format, v...)
}

// Log.Debug("begin listen :8080")
// Log.Debugf("begin listen %d", 8080)
// Log.Info("begin listen :8080")
// Log.Infof("begin listen %d", 8080)
// Log.Warn("begin listen :8080")
// Log.Warnf("begin listen %d", 8080)
// Log.Error("begin listen :8080")
// Log.Errorf("begin listen %d", 8080)
