package logger

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

var (
	logDir string
)

func setup() {
	str, err := os.Getwd()
	if err != nil {
		return
	}

	logDir = str + "/../logs"

	log := GetInstance()
	log.SetDir(logDir)
}

func teardown() {
	os.RemoveAll(logDir)
}

func TestFileLevelNo(t *testing.T) {
	GetInstance().SetLevel(LOGGER_NO)
	GetInstance().Info("test")
	GetInstance().Debug("test")
	GetInstance().Warning("test")
	GetInstance().Error("test")
	logFile := logDir + "/" + GetInstance().fileName
	if !IsFile(logFile) {
		t.Fatalf("test fail, %s is not exists", logFile)
	}

	content, err := ioutil.ReadFile(logFile)
	if err != nil {
		t.Fatalf("read content fail, error: %s", err)
	}

	if len(content) > 0 {
		t.Fatalf("test fail, content is not empty: %s", string(content))
	}
}

func TestFileLevelInfo(t *testing.T) {
	GetInstance().SetLevel(LOGGER_INFO)
	GetInstance().Info("test")
	GetInstance().Debug("test")
	GetInstance().Warning("test")
	GetInstance().Error("test")
	logFile := logDir + "/" + GetInstance().fileName
	if !IsFile(logFile) {
		t.Fatalf("test fail, %s is not exists", logFile)
	}

	content, err := ioutil.ReadFile(logFile)
	if err != nil {
		t.Fatalf("read content fail, error: %s", err)
	}

	con := string(content)
	t.Logf("logs: %s", con)

	strs := strings.Split(con, "\n")
	if len(strs) != 5 {
		t.Fatalf("test fail, len: %d", len(strs))
	}
}

func TestFileLevelDebug(t *testing.T) {
	GetInstance().SetLevel(LOGGER_DEBUG)
	GetInstance().Info("test")
	GetInstance().Debug("test")
	GetInstance().Warning("test")
	GetInstance().Error("test")
	logFile := logDir + "/" + GetInstance().fileName
	if !IsFile(logFile) {
		t.Fatalf("test fail, %s is not exists", logFile)
	}

	content, err := ioutil.ReadFile(logFile)
	if err != nil {
		t.Fatalf("read content fail, error: %s", err)
	}

	con := string(content)
	t.Logf("logs: %s", con)

	strs := strings.Split(con, "\n")
	if len(strs) != 8 {
		t.Fatalf("test fail, len: %d", len(strs))
	}
}

func TestFileLevelWarning(t *testing.T) {
	GetInstance().SetLevel(LOGGER_WARNING)
	GetInstance().Info("test")
	GetInstance().Debug("test")
	GetInstance().Warning("test")
	GetInstance().Error("test")
	logFile := logDir + "/" + GetInstance().fileName
	if !IsFile(logFile) {
		t.Fatalf("test fail, %s is not exists", logFile)
	}

	content, err := ioutil.ReadFile(logFile)
	if err != nil {
		t.Fatalf("read content fail, error: %s", err)
	}

	con := string(content)
	t.Logf("logs: %s", con)

	strs := strings.Split(con, "\n")
	if len(strs) != 10 {
		t.Fatalf("test fail, len: %d", len(strs))
	}
}

func TestFileLevelError(t *testing.T) {
	GetInstance().SetLevel(LOGGER_ERROR)
	GetInstance().Info("test")
	GetInstance().Debug("test")
	GetInstance().Warning("test")
	GetInstance().Error("test")
	logFile := logDir + "/" + GetInstance().fileName
	if !IsFile(logFile) {
		t.Fatalf("test fail, %s is not exists", logFile)
	}

	content, err := ioutil.ReadFile(logFile)
	if err != nil {
		t.Fatalf("read content fail, error: %s", err)
	}

	con := string(content)
	t.Logf("logs: %s", con)

	strs := strings.Split(con, "\n")
	if len(strs) != 11 {
		t.Fatalf("test fail, len: %d", len(strs))
	}
}

func TestMain(m *testing.M) {
	setup()
	m.Run()
	teardown()
}
