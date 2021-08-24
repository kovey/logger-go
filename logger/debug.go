package logger

import (
	"fmt"
	"runtime"
	"time"
)

const (
	logFormat  = "[%s][%s] %s on %s(%d) \n"
	dateFormat = "2006-01-02 15:04:05"
	caller     = 1
)

var level int

func SetLevel(lv int) {
	level = lv
}

func SetLevelByName(name string) {
	lv, ok := loggerMap[name]
	if !ok {
		return
	}

	level = lv
}

func Info(format string, params ...interface{}) {
	if level > LOGGER_INFO || level == LOGGER_NO {
		return
	}

	_, file, line, _ := runtime.Caller(caller)
	fmt.Printf(logFormat, "info", time.Now().Format(dateFormat), fmt.Sprintf(format, params...), file, line)
}

func Debug(format string, params ...interface{}) {
	if level > LOGGER_DEBUG || level == LOGGER_NO {
		return
	}

	_, file, line, _ := runtime.Caller(caller)
	fmt.Printf(logFormat, "debug", time.Now().Format(dateFormat), fmt.Sprintf(format, params...), file, line)
}

func Warning(format string, params ...interface{}) {
	if level > LOGGER_WARNING || level == LOGGER_NO {
		return
	}

	_, file, line, _ := runtime.Caller(caller)
	fmt.Printf(logFormat, "warning", time.Now().Format(dateFormat), fmt.Sprintf(format, params...), file, line)
}

func Error(format string, params ...interface{}) {
	if level == LOGGER_NO {
		return
	}

	_, file, line, _ := runtime.Caller(caller)
	fmt.Printf(logFormat, "error", time.Now().Format(dateFormat), fmt.Sprintf(format, params...), file, line)
}

func Panic(err interface{}) bool {
	if err == nil {
		Debug("err is nil")
		return false
	}

	Error("panic error[%s]", err)

	for i := 3; ; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		Error("%s(%d)", file, line)
	}

	return true
}
