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

func eraseTestFile() {
	dirname, _ := os.Getwd()
	_ = os.Remove(path.Join(dirname, testFileName))
}

func TestNewReporter(t *testing.T) {
	defer eraseTestFile()

	is := is2.New(t)
	cfg := &config.Settings{
		Host:     "localhost",
		Port:     "1234",
		Clients:  "5",
		FileName: testFileName,
	}
	rep := NewReporter(cfg)
	//add some content via the reporter
	err := rep.Append("this is a test")
	is.NoErr(err)

	//check if content was properly saved.
	dirname, _ := os.Getwd()
	raw, err := os.ReadFile(path.Join(dirname, testFileName))
	data := strings.TrimSpace(string(raw))
	is.NoErr(err)

	is.True(data == "this is a test")
}

func TestNewReporter_WithDuplicatedValue(t *testing.T) {
	defer eraseTestFile()

	is := is2.New(t)
	cfg := &config.Settings{
		Host:     "localhost",
		Port:     "1234",
		Clients:  "5",
		FileName: testFileName,
	}
	rep := NewReporter(cfg)

	//add some content via the reporter
	err := rep.Append("this is a test")
	is.NoErr(err)

	err = rep.Append("this is a test")
	is.True(err != nil) //duplication error
	is.Equal(err.Error(), "cannot append duplicate value")

	//check if content was properly saved.
	dirname, _ := os.Getwd()
	raw, err := os.ReadFile(path.Join(dirname, testFileName))
	data := strings.TrimSpace(string(raw))
	is.NoErr(err)

	is.True(data == "this is a test")
}
