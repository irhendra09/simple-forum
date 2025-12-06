package handler

import (
	"net/http"

	"donedev.com/simple-forum/internal/model"
	"donedev.com/simple-forum/internal/service"
	"github.com/gin-gonic/gin"
)

func CreatePostsHandler(c *gin.Context) {
	var request model.CreatePostRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		return
	}
	userId := c.GetInt64("userId")
	err := service.CreatePost(&request, userId)
	if err != nil {
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "successfully created post",
	})
}
