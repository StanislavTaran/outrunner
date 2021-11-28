package main

import (
	"connector/internal/server"
	"log"
)

func main() {
	cfg := server.NewConfig()

	s := server.New(cfg)

	err := s.Run()
	if err != nil {
		log.Fatal("Server has not started... ", err)
	}
}
