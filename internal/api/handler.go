package handlers

import (
	"bufio"
	"fmt"
	"net"
	"petegabriel/central-concurrent-log/pkg/config"
	"strconv"
	"strings"
)

const (
	TerminateCmd = "TERMINATE"
	TerminateCmdMsg = "Terminating process.\nClosing connection.\n"
)


//HandleNewClient receive a new connection to a client and a channel
//where it can check if client can connect to the application or if
//connection need to be released due to overflow of clients.
func HandleNewClient(c net.Conn, sem chan int, st *config.Settings, end chan bool) {

	amount, err := strconv.Atoi(st.Clients)
	if err != nil {
		panic(err)
	}
	if len(sem) == amount {
		c.Write([]byte("Cannot accept more clients.."))
		c.Close()
		return
	}
	sem <- 1

	for {
		cmd, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println("Error communicating with client. Closing connection.")
			c.Close()
			<-sem
			return
		}
		//handle terminate cmd
		if strings.TrimSpace(cmd) == TerminateCmd {
			c.Write([]byte(TerminateCmdMsg))
			c.Close()
			<-sem
			end <- true
			return
		}

		c.Write([]byte("Hello from TCP server\n"))
	}


}
