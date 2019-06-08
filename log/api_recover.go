package log

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"path"
	"strings"
	"github.com/Viktor19931/books_api/utils"
)

func AppRecover(i interface{}, err *error) {
	if r := recover(); r != nil {
		msg := fmt.Sprintf("PANIC %v", r)
		Error(strings.Repeat("-", len(msg)))
		Error(msg)
		Error("> CALL STACK:")
		for _, msg := range logStack() {
			Error("> %s", msg)
		}
		Error(strings.Repeat("-", len(msg)))
		if err != nil {
			*err = errors.New(fmt.Sprintf("PANIC:%+v", r))
		}
	}
}

func logStack() (msgs []string) {
	i := 4
	for {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
        // make file path when was error
		pathElements := strings.Split(file, "/"/*string(os.PathSeparator)*/)
		index := utils.SliceIndex(pathElements, "src")
		if index >= 0 {
			pathElements = pathElements[index+1:]
		}
		file = path.Join(pathElements...)

		msgs = append(msgs, fmt.Sprintf("%s --> %s(%d)", file, path.Base(runtime.FuncForPC(pc).Name()), line))
		i++
		if file == path.Join(pathElements[0], "main.go") {
			break
		}
	}
	return
}

func getFunctionName(i interface{}) string {
	f := runtime.FuncForPC(reflect.ValueOf(i).Pointer())
	if f == nil {
		return ""
	}
	file, line := f.FileLine(f.Entry())
	return fmt.Sprintf("%s:%d %s()", file, line, f.Name())
}

