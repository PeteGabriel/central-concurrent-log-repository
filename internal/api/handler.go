package handlers

import (
	"github.com/rs/zerolog/log"
	"petegabriel/central-concurrent-log/pkg/services"
	"regexp"
)

const (
	digitsRegexPattern = `^\d{9}$`
	terminateCmd       = "TERMINATE"
)

//HandleNewClient receive a new connection to a client and a channel
//where it can check if client can connect to the application or if
//connection need to be released due to overflow of clients.
func HandleNewClient(messenger services.IMessenger, reporter services.IReporter, sem chan int,  end chan bool) {
	sem <- 1 //let other know a new client is in

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
		if cmd == terminateCmd {
			if err := messenger.SendAndTerminate(); err != nil {
				log.Error().Err(err).Msg("Error sending 'terminate message to client.")
			}
			//<-sem
			end <- true
			return
		}

		//check 9 digit condition
		if match, _ := regexp.MatchString(digitsRegexPattern, cmd); !match {
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

		if err := messenger.Send("!! Good input !!"); err != nil {
			log.Error().Err(err).Msg("Error communicating with client.")
		} else {
		  reporter.Append(cmd)
		}
	}
}

