package main

import (
	"connector/internal/server"
	"flag"
	"log"
)

var (
	serverConfigPath string = "config/server.json"
)

func init() {
	flag.StringVar(&serverConfigPath, "server-config-path", "config/server.json", "path to config file")
}

func main() {
	cfg := server.NewConfig()
	if err := server.ReadConfig(serverConfigPath, cfg); err != nil {
		log.Fatal("Could not read config from ", serverConfigPath)
	}
	s := server.New(cfg)

	err := s.Run()
	if err != nil {
		log.Fatal("Server has not started... ", err)
	}
}
