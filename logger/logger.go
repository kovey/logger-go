package logger

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

const (
	LOG_FORMAT = "[%s] %s\n"
)

const (
	LOGGER_NO = iota
	LOGGER_INFO
	LOGGER_DEBUG
	LOGGER_WARNING
	LOGGER_ERROR
)

var (
	loggerMap map[string]int
)

type Logger struct {
	logDir        string
	logDateFormat string
	fileName      string
	logger        *log.Logger
	logLevel      int
	needPrefix    bool
}

var instance *Logger
var once sync.Once

func GetInstance() *Logger {
	once.Do(func() {
		instance = &Logger{
			logDir: "", logDateFormat: "YYYY-MM-DD HH:ii:ss", fileName: time.Now().Format("2006-01-02") + ".log", logLevel: LOGGER_NO, needPrefix: true,
		}
	})

	return instance
}

func NewLogger(logDir string) *Logger {
	log := &Logger{
		logDir: logDir, logDateFormat: "YYYY-MM-DD HH:ii:ss", fileName: time.Now().Format("2006-01-02") + ".log", logLevel: LOGGER_INFO, needPrefix: false,
	}
	return log.SetDir(logDir)
}

func (l *Logger) SetDir(dir string) *Logger {
	if IsFile(dir) {
		panic(fmt.Sprintf("%s is a file\n", dir))
	}

	if !IsDir(dir) {
		os.MkdirAll(dir, os.ModePerm)
	}

	l.logDir = dir
	return l.setLogFile(false)
}

func (l *Logger) setLogFile(check bool) *Logger {
	if check {
		fileName := time.Now().Format("2006-01-02") + ".log"
		if l.fileName == fileName {
			return l
		}

		l.fileName = fileName
	}

	logFile, err := os.OpenFile(l.logDir+"/"+l.fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		Error("open log file error: %s", err)
		return l
	}

	var flag int = 0
	if l.needPrefix {
		flag = log.Ldate | log.Ltime | log.Lshortfile
	}

	l.logger = log.New(logFile, "", flag)
	return l
}

func (l *Logger) SetLevelByName(name string) *Logger {
	lv, ok := loggerMap[name]
	if !ok {
		return l
	}

	l.logLevel = lv
	return l
}

func (l *Logger) SetLevel(level int) *Logger {
	l.logLevel = level
	return l
}

func (l *Logger) Info(format string, v ...interface{}) {
	if l.logLevel > LOGGER_INFO || l.logLevel == LOGGER_NO {
		return
	}

	l.setLogFile(true)
	l.logger.Printf(LOG_FORMAT, "info", fmt.Sprintf(format, v...))
}

func (l *Logger) Debug(format string, v ...interface{}) {
	if l.logLevel > LOGGER_DEBUG || l.logLevel == LOGGER_NO {
		return
	}

	l.setLogFile(true)
	l.logger.Printf(LOG_FORMAT, "debug", fmt.Sprintf(format, v...))
}

func (l *Logger) Warning(format string, v ...interface{}) {
	if l.logLevel > LOGGER_WARNING || l.logLevel == LOGGER_NO {
		return
	}

	l.setLogFile(true)
	l.logger.Printf(LOG_FORMAT, "warning", fmt.Sprintf(format, v...))
}

func (l *Logger) Error(format string, v ...interface{}) {
	if l.logLevel == LOGGER_NO {
		return
	}

	l.setLogFile(true)
	l.logger.Printf(LOG_FORMAT, "error", fmt.Sprintf(format, v...))
}

func (l *Logger) Write(content string) {
	l.setLogFile(true)
	l.logger.Println(content)
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}

	return true
}

func IsDir(path string) bool {
	p, err := os.Stat(path)
	if err != nil {
		return false
	}

	return p.IsDir()
}

func IsFile(path string) bool {
	p, err := os.Stat(path)
	if err != nil {
		return false
	}

	return !p.IsDir()
}

func (l *Logger) GetFileName() string {
	return l.fileName
}

func init() {
	loggerMap = make(map[string]int)

	loggerMap["LOGGER_NO"] = LOGGER_NO
	loggerMap["LOGGER_INFO"] = LOGGER_INFO
	loggerMap["LOGGER_DEBUG"] = LOGGER_DEBUG
	loggerMap["LOGGER_WARNING"] = LOGGER_WARNING
	loggerMap["LOGGER_ERROR"] = LOGGER_ERROR
}
