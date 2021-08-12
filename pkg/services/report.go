package services

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"petegabriel/central-concurrent-log/pkg/config"
)

type IReporter interface {
	Append(n string)
}

type Reporter struct {
	Settings *config.Settings
	writer *bufio.Writer
}

func NewReporter(st *config.Settings) IReporter{
	dirname, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Current directory: %v\n", dirname)
	f, err := os.Create(path.Join(dirname, st.FileName))
	if err != nil {
		panic(err)
	}

	return &Reporter{
		Settings: st,
		writer: bufio.NewWriter(f),
	}
}

func (r *Reporter) Append(n string){
	r.writer.WriteString(n + "\n")
	r.writer.Flush()
}
