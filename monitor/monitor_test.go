package monitor

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
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
	data := make(map[string]interface{})
	data["delay"] = 0.47
	data["class"] = "Hello"
	data["method"] = "World"
	data["time"] = time.Now().Unix()

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
