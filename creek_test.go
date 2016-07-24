package creek

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestLog(t *testing.T) {
	timeNano := strconv.FormatInt(time.Now().UnixNano(), 10)
	path := "testdir_" + timeNano

	logger := log.New(New(path+"/test.log", 1), "Creek Test: ", 0)
	logger.Println("Test line 1")
	logger.Println("Test line 2")

	file, err := os.Open(path + "/test.log")
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	var fileLogs []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileLogs = append(fileLogs, scanner.Text())
	}

	if fileLogs[0] != "Creek Test: Test line 1" {
		t.Errorf("Expected first log to be \"Creek Test: Test line 1\", got %s", fileLogs[0])
	}
	if fileLogs[1] != "Creek Test: Test line 2" {
		t.Errorf("Expected first log to be \"Creek Test: Test line 2\", got %s", fileLogs[1])
	}

	if err := os.RemoveAll(path); err != nil {
		t.Error(err)
	}
}

func TestRollover(t *testing.T) {
	timeNano := strconv.FormatInt(time.Now().UnixNano(), 10)
	path := "testdir_" + timeNano
	megabyte := int64(1024 * 1024)

	logger := log.New(New(path+"/test.log", 1), "", 0)

	for i := 0; int64(i) < megabyte/2; i++ {
		logger.Print("a")
	}

	file, err := os.Open(path + "/test.log")
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		t.Error(err)
	}

	if info.Size() != megabyte {
		t.Errorf("Expected file size to be %d, got %d", megabyte, info.Size())
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		t.Error(err)
	}

	if len(files) != 1 {
		t.Errorf("Expected only one file, found %d", len(files))
	}

	logger.Print("b")

	files, err = ioutil.ReadDir(path)
	if err != nil {
		t.Error(err)
	}

	if len(files) != 2 {
		t.Errorf("Expected to find two files, found %d", len(files))
	}

	if err := os.RemoveAll(path); err != nil {
		t.Error(err)
	}
}
