package main

import (
	"flag"
	"github.com/StanislavTaran/outrunner/internal/server"
	"log"
)

var (
	serverConfigPath = "config/server.json"
)

func init() {
	flag.StringVar(&serverConfigPath, "server-config-path", "./config/server.json", "path to config file")
}

func main() {
	cfg := server.NewConfig()
	if err := server.ReadConfig(serverConfigPath, cfg); err != nil {
		log.Fatal("Could not read config from " + serverConfigPath + "... " + err.Error())
	}
	s := server.New(cfg)

	err := s.Run()
	if err != nil {
		log.Fatal("Server has not started... ", err.Error())
	}
}
