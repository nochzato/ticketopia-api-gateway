package server

import (
	"github.com/gin-gonic/gin"
)

// Server serves the http requests.
type Server struct {
	router *gin.Engine
}

// NewServer creates a new server.
func NewServer() (*Server, error) {
	server := &Server{}

	server.setupRouter()

	return server, nil
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}
