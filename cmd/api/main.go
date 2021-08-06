package main

import (
	"fmt"
	"log"
	"net"
	handlers "petegabriel/central-concurrent-log/internal/api"
	"petegabriel/central-concurrent-log/pkg/config"
	"strconv"
)

func main() {
	s := config.New()
	l, err := net.Listen("tcp", s.Host+":"+s.Port)
	if err != nil {
		panic(err)
	}

	fmt.Println(l.Addr().String())

	defer func(l net.Listener) {
		err := l.Close()
		if err != nil {
			panic(err)
		}
	}(l)

	amnt, err := strconv.Atoi(s.Clients)
	if err != nil {
		log.Println("Accepting 5 clients")
		amnt = 5 // default value of accepted clients
	}
	//use this buffered channel to control how many
	//clients can connect to us simultaneously
	sem := make(chan int, amnt)

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handlers.HandleNewClient(c, sem, s)
	}
}
