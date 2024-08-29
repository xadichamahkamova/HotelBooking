package handler

import (
	pb "api-gateway/genproto/userpb"
	"api-gateway/internal/http/token"
	"api-gateway/logger"

	"github.com/gin-gonic/gin"
)

// @Router /api/users [post]
// @Summary REGISTER USERS
// @Security  		BearerAuth
// @Description This method registers a new user
// @Tags AUTH
// @Accept json
// @Produce json
// @Param user body models.RegisterUserRequest true "User"
// @Success 200 {object} models.RegisterUserResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) RegisterUser(c *gin.Context) {

	req := pb.RegisterUserRequest{}
	if err := c.BindJSON(&req); err != nil {
		logger.Error("Error binding request: ", err)
		c.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
	}
	resp, err := h.Service.RegisterUser(ctx, &req)
	if err != nil {
		logger.Error("Error registering user: ", err)
		c.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
	}
	logger.Info("User registered successfully")
	c.JSON(200, resp)
}

// @Router /api/users/login [post]
// @Summary LOGIN USERS
// @Security  		BearerAuth
// @Description This method login users
// @Security BearerAuth
// @Tags AUTH
// @Accept json
// @Produce json
// @Param user body models.LoginUserRequest true "User"
// @Success 200 {object} token.Tokens
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) LoginUser(c *gin.Context) {

	req := pb.LoginUserRequest{}
	if err := c.BindJSON(&req); err != nil {
		logger.Error("Error binding request: ", err)	
		c.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
	}
	resp, err := h.Service.LoginUser(ctx, &req)
	if err != nil {
		logger.Error("Error logging in user: ", err)
		c.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
	}
	token := token.GenereteJWTToken(resp)
	logger.Info("User logged in successfully")
	c.JSON(200, token)
}
