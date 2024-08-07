package api

import (
	"net/http"

	userPb "github.com/nochzato/ticketopia-user-service/pkg/pb/user/v1"

	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

func (s *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	grpcReq := &userPb.CreateUserRequest{
		Username: req.Username,
		Password: req.Password,
		FullName: req.FullName,
		Email:    req.Email,
	}

	res, err := s.userServiceClient.CreateUser(ctx, grpcReq)
	// TODO: check grpc error code and return right http code
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, res)
}
