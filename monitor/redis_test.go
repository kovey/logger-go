package monitor

import (
	"io/ioutil"
	"os"
	"testing"
)

var (
	redisDir string
)

func setupRedis() {
	str, err := os.Getwd()
	if err != nil {
		return
	}

	redisDir = str + "/../logs/redis"
	InitRedis(redisDir)
}

func teardownRedis() {
	os.RemoveAll(redisDir)
}

func TestRedisWrite(t *testing.T) {
	data := Redis{
		Db:        1,
		Type:      "master",
		Host:      "127.0.0.1",
		Port:      6379,
		Delay:     0.45,
		Command:   "hGet",
		Args:      []interface{}{"key", 1},
		Time:      12324234,
		Timestamp: "2021-01-01 11:11:11",
		Minute:    4353453,
		Date:      "2021-01-01",
		Result:    true,
	}

	data.Write()

	logFile := redisDir + "/" + redisLog.GetFileName()
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
