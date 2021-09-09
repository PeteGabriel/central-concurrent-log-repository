package services

import (
	"os"
	"path"
	"petegabriel/central-concurrent-log/pkg/config"
	"strings"
	"testing"

	is2 "github.com/matryer/is"
)

const testFileName = "number_test.log"

func setupTest() func() {
	// Setup code here

	// tear down later
	return func() {
		// tear-down code here
		dirname, _ := os.Getwd()
		_ = os.Remove(path.Join(dirname, testFileName))
	}
}

func TestNewReporter(t *testing.T) {
	defer setupTest()()

	is := is2.New(t)
	cfg := &config.Settings{
		Host:     "localhost",
		Port:     "1234",
		Clients:  "5",
		FileName: testFileName,
	}
	rep := NewReporter(cfg)
	//add some content via the reporter
	rep.Append("this is a test")

	//check if content was properly saved.
	dirname, _ := os.Getwd()
	raw, err := os.ReadFile(path.Join(dirname, testFileName))
	data := strings.TrimSpace(string(raw))
	is.NoErr(err)

	is.True(data == "this is a test")
}