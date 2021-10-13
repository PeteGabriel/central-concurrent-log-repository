package main

import (
	"net"
	handlers "petegabriel/central-concurrent-log/internal/api"
	"petegabriel/central-concurrent-log/pkg/config"
	"petegabriel/central-concurrent-log/pkg/services"
	"strconv"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	s := config.New()
	l, err := net.Listen("tcp", s.Host+":"+s.Port)
	if err != nil {
		panic(err)
	}

	log.Info().Msg(l.Addr().String())

	defer func(l net.Listener) {
		err := l.Close()
		if err != nil {
			panic(err)
		}
		log.Info().Msg("All connections closed")
	}(l)

	amnt, err := strconv.Atoi(s.Clients)
	if err != nil {
		log.Info().Msg("Accepting 5 clients")
		amnt = 5 // default value of accepted clients
	}
	//use this buffered channel to control how many
	//clients can connect to us simultaneously
	sem := make(chan int, amnt)
	//use this channel to communicate between the goroutine responsible for each client and the main thread.
	//Whenever a client sends the TERMINATE cmd, this channel "allows" the program to finish.
	terminator := make(chan bool)

	go startClientHandler(l, s, sem, terminator)

	<-terminator
}

func startClientHandler(l net.Listener, s *config.Settings, sem chan int, terminator chan bool) {
	for {
		c, err := l.Accept()
		if err != nil {
			log.Error().Err(err).Msg("error accepting client connection")
			return
		}
		msgr := services.New(c)

		amount, err := strconv.Atoi(s.Clients)
		if err != nil {
			panic(err)
		}
		if len(sem) == amount {
			
			err := msgr.Send("Cannot accept more clients..")
			if err != nil {
				log.Error().Err(err).Msg("error sending message to client")
			}

			err = msgr.SendAndTerminate()
			if err != nil {
				log.Error().Err(err).Msg("error sending 'terminate' message to client")
			}
			
			return
		}

		rptr := services.NewReporter(s)

		go handlers.HandleNewClient(msgr, rptr, sem, terminator)
	}
}
