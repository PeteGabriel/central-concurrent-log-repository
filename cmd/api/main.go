package main

import (
	"fmt"
	"net"
	"petegabriel/central-concurrent-log/pkg/config"
)

func main() {
	s := config.New()
	l, err := net.Listen("tcp", s.Host+s.Port)
	if err != nil {
		panic(err)
	}

	fmt.Println(l.Addr().String())

	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Found a stranger")
		c.Write([]byte("Hello stranger"))
	}
}
