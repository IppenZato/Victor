package log

import (
	"github.com/sirupsen/logrus"
	"strings"
	"time"
	"path"
	"runtime"
	"fmt"
	"github.com/Viktor19931/books_api/utils"
)

const (
	locationSkip       = 6
	locationTag        = "%loc%"
	// Default log format will output [INFO]: 2006-01-02T15:04:05Z07:00 - Log message
	defaultLogFormat       = "[%lvl%]: %time% - %msg%"
	defaultTimestampFormat = time.RFC3339
)

// formatter implements logrus.formatter interface.
type formatter struct {
	// Timestamp format
	TimestampFormat string
	// Available standard keys: time, msg, lvl
	// Also can include custom fields but limited to strings.
	// All of fields need to be wrapped inside %% i.e %time% %msg%
	LogFormat string
}

// Format building log message.
func (f *formatter) Format(entry *logrus.Entry) ([]byte, error) {
	output := f.LogFormat
	if output == "" {
		output = defaultLogFormat
	}

	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = defaultTimestampFormat
	}

	output = strings.Replace(output, "%time%", entry.Time.Format(timestampFormat), 1)

	level := strings.ToUpper(entry.Level.String())[0:1]
	output = strings.Replace(output, "%lvl%", level, 1)

	if strings.Contains(output, locationTag) {
		output = strings.Replace(output, locationTag, f.when(locationSkip), 1)
	}

	output = strings.Replace(output, "%msg%", entry.Message, 1)

	for k, v := range entry.Data {
		if s, ok := v.(string); ok {
			output = strings.Replace(output, "%"+k+"%", s, 1)
		}
	}

	return []byte(output), nil
}

func (f *formatter) when(skip int) string {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "unknown:0:unknown()"
	}

	// make file path when was error
	pathElements := strings.Split(file, "/"/*string(os.PathSeparator)*/)
	index := utils.SliceIndex(pathElements, "src")
	if index >= 0 {
		pathElements = pathElements[index+2:]
	}
	file = path.Join(pathElements...)
	_ = pc

	//return fmt.Sprintf("%s:%s(%d)", file, path.Base(runtime.FuncForPC(pc).Name()), line)
	return fmt.Sprintf("%s:%d", file, line)
}