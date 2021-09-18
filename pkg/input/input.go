package input

import (
	"bufio"
	"fmt"
	"go-msgr/pkg/display"
	"go-msgr/pkg/transport"
	"os"
)

type Input struct {
	sender      transport.Sender
	disp        *display.Display
	destination string
}

func New(sender transport.Sender, disp *display.Display, destination string) Input {
	return Input{sender, disp, destination}
}

func (i *Input) Run() {
	reader := bufio.NewReader(os.Stdin)
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			i.disp.Show(fmt.Errorf("Input read string error: %w", err).Error())
			continue
		}
		err = i.sender.Send(i.destination, text)
		if err != nil {
			i.disp.Show(fmt.Errorf("Input send error: %w", err).Error())
			continue
		}
	}
}
