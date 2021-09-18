package main

import (
	"log"
	"go-msgr/pkg/config"
)

const version = "v0.0.1"

func main() {
	log.Printf("msgr %v", version)
	var cfg config.Config
	cfg.Init()
	config.Parse()
	cfg.Dump()
}
