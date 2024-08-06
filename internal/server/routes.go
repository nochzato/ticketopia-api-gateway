package server

import (
	"github.com/gin-gonic/gin"
)

func (s *Server) setupRouter() {
	router := gin.Default()

	s.router = router
}
