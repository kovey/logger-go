package monitor

import (
	"io/ioutil"
	"os"
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
	Init(logDir)
}

func teardown() {
	os.RemoveAll(logDir)
}

func TestWrite(t *testing.T) {
	data := Monitor{
		Delay:         0.45,
		Path:          "/index/get",
		RequestMethod: "post",
		Params:        "parmas",
		RequestTime:   123124324,
		Action:        1000,
		ResAction:     1001,
		Packet:        "",
		Base:          "",
		Fd:            1,
		Type:          "rpc",
		Err:           "err",
		Trace:         "trace",
		Service:       "service",
		ServiceType:   "service_type",
		Class:         "class",
		Method:        "method",
		Args:          []string{"test", "aaa", "bb"},
		Ip:            "127.0.0.1",
		Time:          12324234,
		Timestamp:     "2021-01-01 11:11:11",
		Minute:        4353453,
		Response:      true,
		TraceId:       "aaaa",
		From:          "golang",
		End:           23534534,
		ParentId:      "sdfsfsfs",
		SpanId:        "afsadad",
		ClientVersion: "1.0",
		ServerVersion: "1.0",
		ClientLang:    "php",
		ServerLang:    "golang",
		Header:        "header",
	}

	Write(data)
	Write(data)

	logFile := logDir + "/" + log.GetFileName()
	content, err := ioutil.ReadFile(logFile)
	if err != nil {
		t.Fatalf("read content fail, error: %s", err)
	}

	if len(content) == 0 {
		t.Fatalf("test failure, content is empty")
	}

	con := string(content)
	t.Logf("logs: %s", con)
}

func TestMain(m *testing.M) {
	setup()
	m.Run()
	teardown()
}
