package main

import (
	"fmt"
	"log"

	userPb "github.com/nochzato/ticketopia-user-service/pkg/pb/user/v1"

	"github.com/nochzato/ticketopia-api-gateway/internal/api"
	"github.com/nochzato/ticketopia-api-gateway/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot read config file: %s", err)
	}

	log.Println("starting user service grpc connection")
	userServiceConn, err := grpc.NewClient(
		config.UserService.Addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to establish connection with user service: %s", err)
	}
	defer userServiceConn.Close()

	userServiceClient := userPb.NewUserServiceClient(userServiceConn)

	server, err := api.NewServer(userServiceClient)
	if err != nil {
		log.Fatalf("failed to create server")
	}

	httpAddr := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)

	if err = server.Start(httpAddr); err != nil {
		log.Fatalf("failed to start server: %s", err)
	}
}
