package services

import (
	"bufio"
	"net"
	"strings"
)

const terminateCmdMsg = "Terminating process...Closing connection...\n"

//IMessenger specifies the operations
type IMessenger interface {
	SendAndTerminate() error
	Send(msg string) error
	Read() (string, error)
	CloseSession() error
}

type Messenger struct {
	c net.Conn
}

//New instance of Messenger
func New(c net.Conn) IMessenger {
	return &Messenger{
		c: c,
	}
}

//SendAndTerminate sends a last message to the client and
//closes the connection.
func (m *Messenger) SendAndTerminate() error {
	if _, err := m.sendMsg(terminateCmdMsg); err != nil {
		return err
	}

	return m.CloseSession()
}

//Send a specific message.
func (m *Messenger) Send(msg string) error {
	if _, err := m.sendMsg(msg); err != nil {
		return err
	}
	return nil
}

//Read a message. Blocking operation.
func (m *Messenger) Read() (string, error) {
	cmd, err := bufio.NewReader(m.c).ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(cmd), nil
}

func (m *Messenger) CloseSession() error {
	if err := m.c.Close(); err != nil {
		return err
	}
	return nil
}

func (m *Messenger) sendMsg(msg string) (n int, err error) {
	if _, err := m.c.Write([]byte(msg + "\n")); err != nil {
		return -1, err
	}

	return len(msg), nil
}