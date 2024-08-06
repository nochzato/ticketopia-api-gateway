package main

import (
	"fmt"
	"log"

	"github.com/nochzato/ticketopia-api-gateway/internal/config"
	"github.com/nochzato/ticketopia-api-gateway/internal/server"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot read config file: %s", err)
	}

	server, err := server.NewServer()
	if err != nil {
		log.Fatalf("failed to create server")
	}

	httpAddr := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)

	err = server.Start(httpAddr)
	if err != nil {
		log.Fatalf("failed to start server: %s", err)
	}
}
