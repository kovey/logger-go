package monitor

import (
	"encoding/json"

	"github.com/kovey/logger-go/logger"
)

var (
	log *logger.Logger
)

func Init(logDir string) {
	log = logger.NewLogger(logDir)
}

func Write(data map[string]interface{}) {
	content, err := json.Marshal(data)
	if err != nil {
		logger.Error("monitor write failure, error[%s]", err)
		return
	}

	log.Write(string(content))
}
