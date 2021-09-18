package main

import (
	"fmt"
	"go-msgr/pkg/config"
	"go-msgr/pkg/display"
)

const version = "v0.0.1"

func main() {
	var cfg config.Config
	cfg.Init()
	config.Parse()
	cfg.Dump()
	var disp display.Display
	disp.Show(fmt.Sprintf("msgr %v", version))
}
