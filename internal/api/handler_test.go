package handlers

import (
	is2 "github.com/matryer/is"
	"petegabriel/central-concurrent-log/pkg/config"
	"petegabriel/central-concurrent-log/pkg/mocks"
	"testing"
)

func TestHandleNewClient_Ok(t *testing.T) {
	is := is2.New(t)
	settings := config.Settings{
		Host:    "localhost",
		Port:    "4001",
		Clients: "5",
	}
	sem := make(chan int)
	end := make(chan bool)
	mockedConn := mocks.NewClient()

	HandleNewClient(mockedConn, sem, &settings, end)

	is.Equal(len(sem), 1) //semaphore length should be 1
	is.Equal(len(end), 0) //we do not want end channel to have unread elements in it

	is.Equal(mockedConn.Container[0], "Hello from TCP server\n")
}