package config

import (
	"flag"
	"log"
)

const (
	listenAddrKey = "ListenAddr"
	destinationAddrKey = "DestinationAddr"
	filePathKey = "FilePath"
)

type Config struct {
	ListenAddr string
	DestinationAddr string
	FilePath string
}

func (c *Config) Init() {
	flag.StringVar(&c.ListenAddr, listenAddrKey, "0.0.0.0:8080", "Address to listen messages on.")
	flag.StringVar(&c.DestinationAddr, destinationAddrKey, "0.0.0.0:8080", "Address to send messages to.")
	flag.StringVar(&c.FilePath, filePathKey, "./config.json", "Path to config file.")
}

func dumpImpl(key, value string) {
	log.Printf("%s: %s", key, value)
}

func (c Config) Dump() {
	dumpImpl(listenAddrKey, c.ListenAddr)
	dumpImpl(destinationAddrKey, c.DestinationAddr)
	dumpImpl(filePathKey, c.FilePath)
}

var Parse = flag.Parse
