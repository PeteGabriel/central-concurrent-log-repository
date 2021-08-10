package handlers

import (
	"bufio"
	"fmt"
	"net"
	"petegabriel/central-concurrent-log/pkg/config"
	"petegabriel/central-concurrent-log/pkg/services"
	"strconv"
	"strings"
)

const terminateCmd = "TERMINATE"


//HandleNewClient receive a new connection to a client and a channel
//where it can check if client can connect to the application or if
//connection need to be released due to overflow of clients.
func HandleNewClient(c net.Conn, sem chan int, st *config.Settings, end chan bool) {

	amount, err := strconv.Atoi(st.Clients)
	if err != nil {
		panic(err)
	}
	if len(sem) == amount {
		services.SendMsg(c, "Cannot accept more clients..")
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
		if strings.TrimSpace(cmd) == terminateCmd {
			services.SendTerminatorConsoleMsg(c)
			<-sem
			end <- true
			return
		}

		services.SendMsg(c, "Hello from TCP server\n")
	}


}
