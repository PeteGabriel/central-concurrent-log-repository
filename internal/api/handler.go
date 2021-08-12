package handlers

import (
	"petegabriel/central-concurrent-log/pkg/config"
	"petegabriel/central-concurrent-log/pkg/services"
	"regexp"
	"strconv"
	"github.com/rs/zerolog/log"
)

const (
	DigitsPattern = `^\d{9}$`
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
			log.Error().Err(err).Msg("error sending message to client")
		}
		err = messenger.SendAndTerminate()
		if err != nil {
			log.Error().Err(err).Msg("error sending 'terminate' message to client")
		}
		return
	}
	sem <- 1

	for {
		//accepting input
		cmd, err := messenger.Read()
		if err != nil {
			log.Error().Err(err).Msg("Error communicating with client. Closing connection.")
			_ = messenger.SendAndTerminate()
			<-sem
			return
		}

		//log input for record
		log.Info().Msg(cmd)

		//handle terminate cmd
		if cmd == TerminateCmd {
			if err := messenger.SendAndTerminate(); err != nil {
				log.Error().Err(err).Msg("Error sending 'terminate message to client.")
			}
			//<-sem
			end <- true
			return
		}

		//check 9 digit condition
		if match, _ := regexp.MatchString(DigitsPattern, cmd); !match {
			err := messenger.Send("--> Input length not valid. <-")
			if err != nil {
				log.Error().Err(err).Msg("Error communicating with client.")
			}
			if err := messenger.CloseSession(); err != nil {
				log.Error().Err(err).Msg("")
			}
			<-sem //free space for another client
			return
		}

		//handle client input
		_, err = strconv.Atoi(cmd)
		if err != nil {
			log.Error().Err(err).Msg("could not convert input to number")
			if err := messenger.Send("--> Input not valid <-"); err != nil {
				log.Error().Err(err).Msg("Error communicating with client.")
			}
			if err := messenger.CloseSession(); err != nil {
				log.Error().Err(err).Msg("")
			}
			<-sem //free space for another client
			return
		}

		if err := messenger.Send("!! Good input !!"); err != nil {
			log.Error().Err(err).Msg("Error communicating with client.")
		}

	}
}

