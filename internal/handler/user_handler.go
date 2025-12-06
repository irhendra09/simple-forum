package handler

import (
	"net/http"

	"donedev.com/simple-forum/internal/model"
	"donedev.com/simple-forum/internal/service"
	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var request model.SignUpRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := service.Register(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
		"user_id": user.ID,
	})
}

func SignIn(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := service.Login(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":       "User logged in successfully",
		"access_token":  user.AccessToken,
		"refresh_token": user.RefreshToken,
	})
}
