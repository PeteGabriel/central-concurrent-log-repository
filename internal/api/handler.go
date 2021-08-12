package handlers

import (
	"fmt"
	"petegabriel/central-concurrent-log/pkg/config"
	"petegabriel/central-concurrent-log/pkg/services"
	"strconv"
)

const (
	AmountOfDigits = 9
	TerminateCmd = "TERMINATE"
)



//HandleNewClient receive a new connection to a client and a channel
//where it can check if client can connect to the application or if
//connection need to be released due to overflow of clients.
func HandleNewClient(messenger services.IMessenger, sem chan int, st *config.Settings, end chan bool) {

	amount, err := strconv.Atoi(st.Clients)
	if err != nil {
		panic(err)
	}
	if len(sem) == amount {
		err := messenger.Send("Cannot accept more clients..")
		if err != nil {
			//todo add logs
		}
		err = messenger.SendAndTerminate()
		if err != nil {
			//todo add logs
		}
		return
	}
	sem <- 1

	for {
		//accepting input
		cmd, err := messenger.Read()
		if err != nil {
			fmt.Println("Error communicating with client. Closing connection.")
			err := messenger.SendAndTerminate()
			if err != nil {
				<-sem
				return
			}
			<-sem
			return
		}

		//handle terminate cmd
		if cmd == TerminateCmd {
			terminateProcess(messenger, sem, end)
			return
		}

		//check 9 digit condition
		if len(cmd) != AmountOfDigits {
			err := messenger.Send("--> Input not valid. <-\n")
			if err != nil {
				//todo add logs
			}
			terminateProcess(messenger, sem, end)
			return
		}

		//handle client input
		_, err = strconv.Atoi(cmd)
		if err != nil {
			err := messenger.Send("--> Input not valid <-\n")
			if err != nil {
				//todo add logs
			}
			fmt.Println(err)
		}else {
			fmt.Println("Input received: " + cmd)
			err := messenger.Send("!! Good input !!\n")
			if err != nil {
				//todo add logs
			}
		}
	}
}

func terminateProcess(messenger services.IMessenger, sem chan int, end chan bool) {
	err := messenger.SendAndTerminate()
	if err != nil {
		//todo add logs
		return
	}
	<-sem
	end <- true
}
