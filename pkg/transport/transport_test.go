package transport

import (
	"fmt"
	"go-msgr/pkg/display"
	"testing"
	"time"
)

const testAddress = "localhost:8080"

func TestTransport(t *testing.T) {
	var disp display.Display
	listener := Listener{&disp}
	var sender Sender
	err := listener.Listen(testAddress)
	if err != nil {
		disp.Show(fmt.Errorf("Transport test error: %w", err).Error())
	}
	time.Sleep(1 * time.Second)
	sender.Send(testAddress, "Test message\n")
	time.Sleep(3 * time.Second)
}
