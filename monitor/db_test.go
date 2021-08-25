package monitor

import (
	"io/ioutil"
	"os"
	"testing"
)

var (
	dbDir string
)

func setupDb() {
	str, err := os.Getwd()
	if err != nil {
		return
	}

	dbDir = str + "/../logs/db"
	InitDb(dbDir)
}

func teardownDb() {
	os.RemoveAll(dbDir)
}

func TestDbWrite(t *testing.T) {
	data := Db{
		Dbname:    "test",
		Type:      "master",
		DbType:    "mysql",
		Host:      "127.0.0.1",
		Port:      3306,
		Delay:     0.45,
		Sql:       "select * from user",
		Time:      12324234,
		Timestamp: "2021-01-01 11:11:11",
		Minute:    4353453,
		Date:      "2021-01-01",
		Result:    "12",
	}

	data.Write()

	logFile := dbDir + "/" + dbLog.GetFileName()
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
