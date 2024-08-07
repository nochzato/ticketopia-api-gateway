package api

import (
	"github.com/gin-gonic/gin"
)

func (s *Server) setupRouter() {
	router := gin.Default()

	router.POST("/users", s.createUser)

	s.router = router
}
