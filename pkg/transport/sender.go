package transport

import (
	"net"
	"fmt"
)

type Sender struct {

}

func NewSender() Sender {
	return Sender{}
}

func (s *Sender) Send(destination, message string) error {
	conn, err := net.Dial("tcp", destination)
	if err != nil {
		return fmt.Errorf("Sender dial error: %w", err)
	}
	messageBytes := []byte(message)
	n, err := conn.Write(messageBytes)
	if err != nil {
		return fmt.Errorf("Sender write error: %w", err)
	}
	if n != len(messageBytes) {
		return fmt.Errorf("Sender write error: incomplete message sending")
	}
	return nil
}
