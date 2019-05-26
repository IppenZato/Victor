package log

import (
	"github.com/sirupsen/logrus"
	"strings"
	"fmt"
	"io"
	"os"
	"github.com/Viktor19931/books_api/utils"
)

var logger = openLog()
var logFile *os.File

func openLog() *logrus.Logger {
	log := &logrus.Logger{
		Out:   os.Stderr,
		Level: logrus.DebugLevel,
		Formatter: &formatter{
			TimestampFormat: "2006-01-02 15:04:05.000",
			LogFormat:       "%time% [%lvl%] [%loc%]: %msg%",
		},
	}
	file := utils.ChangeFileExt(utils.RunFile(), ".log")
	logFile, _ = os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	log.Out = io.MultiWriter(os.Stderr, logFile)
	logrus.RegisterExitHandler(closeLogFile)
	return log
}

func closeLogFile() {
	if logFile != nil {
		logFile.Close()
	}
}

func formatLog(f interface{}, v ...interface{}) (msg string) {
	defer func() {
		if len(msg) == 0 || string(msg[len(msg)-1]) != "\n" {
			msg = msg + "\n"
		}
	}()

	switch f.(type) {
	case string:
		msg += f.(string)
		if len(v) == 0 {
			return msg
		}
		if strings.Contains(msg, "%") && !strings.Contains(msg, "%%") {
			//format string
		} else {
			//do not contain format char
			msg += strings.Repeat(" %v", len(v))
		}
	default:
		msg += fmt.Sprint(f)
		if len(v) == 0 {
			return msg
		}
		msg += strings.Repeat(" %v", len(v))
	}

	return fmt.Sprintf(msg, v...)
}
