package api

import (
	"github.com/gin-gonic/gin"
	userPb "github.com/nochzato/ticketopia-user-service/pkg/pb/user/v1"
)

type Server struct {
	router            *gin.Engine
	userServiceClient userPb.UserServiceClient
}

func NewServer(userServiceClient userPb.UserServiceClient) (*Server, error) {
	server := &Server{
		userServiceClient: userServiceClient,
	}

	server.setupRouter()

	return server, nil
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}
