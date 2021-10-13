package services

import (
	"bufio"
	"errors"
	"os"
	"path"
	"petegabriel/central-concurrent-log/pkg/config"
	"petegabriel/central-concurrent-log/pkg/domain"
)

//IReporter specifies the writing operations
type IReporter interface {
	Append(n string) error
}

type Reporter struct {
	Settings *config.Settings
	writer *bufio.Writer
	duplicate domain.IDuplicate
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
		duplicate: domain.New(),
	}
}

//Append new content to report file.
func (r *Reporter) Append(n string) error{
	if r.duplicate.IsDuplicate(n) {
		return errors.New("cannot append duplicate value")
	}

	if _, err := r.writer.WriteString(n + "\n"); err != nil {
		return err
	}else {
		if err := r.duplicate.TrySaveNewValue(n); err != nil {
			return err
		}
	}

	if err := r.writer.Flush(); err != nil {
		return err
	}

	return nil
}
