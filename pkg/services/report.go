package services

import (
	"bufio"
	"os"
	"path"
	"petegabriel/central-concurrent-log/pkg/config"
)

//IReporter specifies the writing operations
type IReporter interface {
	Append(n string)
}

type Reporter struct {
	Settings *config.Settings
	writer *bufio.Writer
}

//NewReporter creates a new instance of Reporter
func NewReporter(st *config.Settings) IReporter{
	dirname, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	f, err := os.Create(path.Join(dirname, st.FileName))
	if err != nil {
		panic(err)
	}

	return &Reporter{
		Settings: st,
		writer: bufio.NewWriter(f),
	}
}

//Append new content to report file.
func (r *Reporter) Append(n string){
        if _, err := r.writer.WriteString(n + "\n"); err != nil {
		return
	}
	if err := r.writer.Flush(); err != nil {
		return
	}
}
