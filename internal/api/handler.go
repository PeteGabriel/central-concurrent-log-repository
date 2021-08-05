package handlers

import (
	"fmt"
	"net"
)

func HandleNewClient(c net.Conn, sem chan int) {

	if len(sem) == 2 {
		c.Write([]byte("Cannot accept more clients.."))
		c.Close()
		return
	}
	sem <- 1
	fmt.Println("Sending msg...")
	c.Write([]byte("Hello from goroutine"))

}
