package handler

import (
	"net/http"

	apperrors "donedev.com/simple-forum/internal/errors"
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
	user, err := service.UserService.Register(&request)
	if err != nil {
		status := apperrors.ToHTTPStatus(err)
		c.JSON(status, gin.H{"error": err.Error()})
		return
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
	user, err := service.UserService.Login(&req)
	if err != nil {
		status := apperrors.ToHTTPStatus(err)
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":       "User logged in successfully",
		"access_token":  user.AccessToken,
		"refresh_token": user.RefreshToken,
	})
}

func RefreshToken(c *gin.Context) {
	var req model.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.RefreshToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing refresh token"})
		return
	}
	resp, err := service.UserService.RefreshToken(req.RefreshToken)
	if err != nil {
		status := apperrors.ToHTTPStatus(err)
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func Logout(c *gin.Context) {
	var req model.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.RefreshToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing refresh token"})
		return
	}
	if err := service.UserService.Logout(req.RefreshToken); err != nil {
		status := apperrors.ToHTTPStatus(err)
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
