package mocks

import (
	"net"
	"time"
)

//Client represents a mocked connection.
//Is just a dummy implementation of the methods
// * Write(b []byte) (n int, err error)
// * Close() error
type Client struct {
	Container []string
}

func (c *Client) Read(b []byte) (n int, err error) {
	panic("implement me")
}

func (c *Client) LocalAddr() net.Addr {
	panic("implement me")
}

func (c *Client) RemoteAddr() net.Addr {
	panic("implement me")
}

func (c *Client) SetDeadline(t time.Time) error {
	panic("implement me")
}

func (c *Client) SetReadDeadline(t time.Time) error {
	panic("implement me")
}

func (c *Client) SetWriteDeadline(t time.Time) error {
	panic("implement me")
}

func NewClient() *Client {
	return &Client{
		Container: make([]string, 10),
	}
}

func (c *Client) Write(b []byte) (n int, err error){
	c.Container = append(c.Container, string(b))
	return len(string(b)), nil
}

func (c *Client) Close() error {
	return nil
}