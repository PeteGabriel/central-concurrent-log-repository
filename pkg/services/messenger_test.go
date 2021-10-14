package services

import (
	"io/ioutil"
	"net"
	"testing"

	is2 "github.com/matryer/is"
)

func TestSend(t *testing.T) {
	is := is2.New(t)

	//Pipe creates a synchronous, in-memory, full duplex network connection
	server, client := net.Pipe()

	m := New(server)

	go func() {
		m.Send("this new message")
		m.CloseSession()
	}()

	//assert that client received the message
	//ioutil.ReadAll function blocks until EOF.
	//The call of close method causes the EOF on the stream.
	buf, err := ioutil.ReadAll(client)
	is.NoErr(err)
	is.Equal(string(buf), "this new message\n")

	client.Close()
}

func TestCloseSession(t *testing.T) {
	is := is2.New(t)

	//Pipe creates a synchronous, in-memory, full duplex network connection
	server, client := net.Pipe()

	m := New(server)

	go func() {
		m.CloseSession()
	}()

	_, err := ioutil.ReadAll(client)
	is.NoErr(err)

	err = m.Send("")
	is.True(err != nil) //connection closed already

	client.Close()

}

func TestSendAndTerminate(t *testing.T) {
	is := is2.New(t)

	server, client := net.Pipe()

	m := New(server)

	go func() {
		m.SendAndTerminate()
	}()

	_, err := ioutil.ReadAll(client)
	is.NoErr(err)

	err = m.Send("")
	is.True(err != nil) //connection closed already

	client.Close()
}

func TestRead(t *testing.T) {
	is := is2.New(t)

	server, client := net.Pipe()

	m := New(server)

	go func() {
		client.Write([]byte("message read"))
		client.Close()
	}()

	content, err := m.Read()
	is.NoErr(err)
	is.Equal(content, "message read")

	server.Close()
}
