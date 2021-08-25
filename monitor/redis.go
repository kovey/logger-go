package monitor

import "github.com/kovey/logger-go/logger"

var (
	redisLog *logger.Logger
)

type Redis struct {
	Db        int           `json:"db"`
	Type      string        `json:"type"`
	Host      string        `json:"host"`
	Port      int           `json:"port"`
	Delay     float64       `json:"delay"`
	Command   string        `json:"sql"`
	Args      []interface{} `json:"args"`
	Time      int64         `json:"time"`
	Timestamp string        `json:"timestamp"`
	Minute    int           `json:"minute"`
	Date      string        `json:"date"`
	Result    interface{}   `json:"result"`
}

func InitRedis(logDir string) {
	redisLog = logger.NewLogger(logDir)
}

func (r Redis) Write() {
	write(r, redisLog)
}
