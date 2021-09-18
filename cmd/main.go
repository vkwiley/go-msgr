package main

import (
	"fmt"
	"go-msgr/pkg/config"
	"go-msgr/pkg/display"
	"go-msgr/pkg/transport"
	"go-msgr/pkg/input"
)

const version = "v0.0.1"

func main() {
	var cfg config.Config
	cfg.Init()
	config.Parse()
	cfg.Dump()
	var disp display.Display
	disp.Show(fmt.Sprintf("msgr %v", version))
	listener := transport.NewListener(&disp)

	err := listener.Listen(cfg.ListenAddr)
	if err != nil {
		disp.Show(fmt.Errorf("Listener error: %w", err).Error())
	}
	sender := transport.NewSender()
	in := input.New(sender, &disp, cfg.DestinationAddr)
	in.Run()
}
