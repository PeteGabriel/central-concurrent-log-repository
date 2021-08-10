package services

import "net"

const terminateCmdMsg = "Terminating process.\nClosing connection.\n"


//SendTerminatorConsoleMsg sends a pre-configured message for when
//the terminate command is received.
func SendTerminatorConsoleMsg(c net.Conn){
	if err := sendMsg(c, terminateCmdMsg); err != nil {
		//TODO add logs
	}
}

//SendMsg send a specified message to the
func SendMsg(c net.Conn, msg string){
	if err := sendMsg(c, msg); err != nil {
		//TODO add logs
	}
}

func sendMsg(c net.Conn, msg string) error{
	if _, err := c.Write([]byte(msg)); err != nil {
		//TODO add logs
		return err
	}
	if err := c.Close(); err != nil {
		//TODO add logs
		return err
	}
	return nil
}