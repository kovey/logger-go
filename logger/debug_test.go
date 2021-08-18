package logger

import "testing"

func TestLevelNo(t *testing.T) {
	SetLevel(LOGGER_NO)
	Info("test info")
	Debug("test debug")
	Warning("test warning")
	Error("test err")
}

func TestLevelInfo(t *testing.T) {
	SetLevel(LOGGER_INFO)
	Info("test info")
	Debug("test debug")
	Warning("test warning")
	Error("test err")
}

func TestLevelDebug(t *testing.T) {
	SetLevel(LOGGER_DEBUG)
	Info("test info")
	Debug("test debug")
	Warning("test warning")
	Error("test err")
}

func TestLevelWarning(t *testing.T) {
	SetLevel(LOGGER_WARNING)
	Info("test info")
	Debug("test debug")
	Warning("test warning")
	Error("test err")
}

func TestLevelError(t *testing.T) {
	SetLevel(LOGGER_ERROR)
	Info("test info")
	Debug("test debug")
	Warning("test warning")
	Error("test err")
}

func TestLevelByNameNo(t *testing.T) {
	SetLevelByName("LOGGER_NO")
	Info("test info")
	Debug("test debug")
	Warning("test warning")
	Error("test err")
}

func TestLevelByNameInfo(t *testing.T) {
	SetLevelByName("LOGGER_INFO")
	Info("test info")
	Debug("test debug")
	Warning("test warning")
	Error("test err")
}

func TestLevelByNameDebug(t *testing.T) {
	SetLevelByName("LOGGER_DEBUG")
	Info("test info")
	Debug("test debug")
	Warning("test warning")
	Error("test err")
}

func TestLevelByNameWarning(t *testing.T) {
	SetLevelByName("LOGGER_WARNING")
	Info("test info")
	Debug("test debug")
	Warning("test warning")
	Error("test err")
}

func TestLevelByNameError(t *testing.T) {
	SetLevelByName("LOGGER_ERROR")
	Info("test info")
	Debug("test debug")
	Warning("test warning")
	Error("test err")
}
