package transport

import (
	"net"
	"fmt"
	"bufio"
	"go-msgr/pkg/display"
)

type Listener struct {
	disp *display.Display
}

func NewListener(disp *display.Display) Listener {
	return Listener{disp}
}

func (l *Listener) Listen(listenAddr string) error {
	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return fmt.Errorf("Listener error: %w", err)
	}
	go l.startAccepting(ln)
	return nil
}

func (l *Listener) startAccepting(ln net.Listener) {
	for {
		conn, err := ln.Accept()
		if err != nil {
			l.disp.Show(fmt.Errorf("Listener accept error: %w", err).Error())
			continue
		}
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			l.disp.Show(fmt.Errorf("Listener read error: %w", err).Error())
			continue
		}
		l.disp.Show(string(message[:len(message)-1]))
	}
}
