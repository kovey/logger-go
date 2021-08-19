package monitor

import (
	"encoding/json"

	"github.com/kovey/logger-go/logger"
)

var (
	log *logger.Logger
)

type Monitor struct {
	Delay         float64     `json:"delay"`
	Path          string      `json:"path"`
	RequestMethod string      `json:"request_method"`
	Params        interface{} `json:"params"`
	RequestTime   int64       `json:"request_time"`
	Action        int         `json:"action"`
	ResAction     int         `json:"res_action"`
	Packet        string      `json:"packet"`
	Base          string      `json:"base"`
	Fd            int         `json:"fd"`
	Type          string      `json:"type"`
	Err           string      `json:"err"`
	Trace         string      `json:"trace"`
	Service       string      `json:"service"`
	ServiceType   string      `json:"service_type"`
	Class         string      `json:"class"`
	Method        string      `json:"method"`
	Args          interface{} `json:"args"`
	Ip            string      `json:"ip"`
	Time          int64       `json:"time"`
	Timestamp     string      `json:"timestamp"`
	Minute        int64       `json:"minute"`
	HttpCode      int         `json:"http_code"`
	Response      interface{} `json:"response"`
	TraceId       string      `json:"traceId"`
	From          string      `json:"from"`
	End           int64       `json:"end"`
	ParentId      string      `json:"parentId"`
	SpandId       string      `json:"spanId"`
	ClientVersion string      `json:"client_version"`
	ServerVersion string      `json:"server_version"`
	ClientLang    string      `json:"client_lang"`
	ServerLang    string      `json:"server_lang"`
	Header        interface{} `json:"header"`
}

func Init(logDir string) {
	log = logger.NewLogger(logDir)
}

func Write(data Monitor) {
	content, err := json.Marshal(data)
	if err != nil {
		logger.Error("monitor write failure, error[%s]", err)
		return
	}

	log.Write(string(content))
}
