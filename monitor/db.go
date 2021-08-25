package monitor

import "github.com/kovey/logger-go/logger"

var (
	dbLog *logger.Logger
)

type Db struct {
	Dbname    string      `json:"dbname"`
	Type      string      `json:"type"`
	DbType    string      `json:"dbType"`
	Host      string      `json:"host"`
	Port      int         `json:"port"`
	Delay     float64     `json:"delay"`
	Sql       string      `json:"sql"`
	Time      int64       `json:"time"`
	Timestamp string      `json:"timestamp"`
	Minute    int         `json:"minute"`
	Date      string      `json:"date"`
	Result    interface{} `json:"result"`
}

func InitDb(logDir string) {
	dbLog = logger.NewLogger(logDir)
}

func (d Db) Write() {
	write(d, dbLog)
}
