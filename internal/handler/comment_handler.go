package handler

import (
	"net/http"
	"strconv"

	apperrors "donedev.com/simple-forum/internal/errors"
	"donedev.com/simple-forum/internal/model"
	"donedev.com/simple-forum/internal/service"
	"github.com/gin-gonic/gin"
)

func CreateCommentHandler(c *gin.Context) {
	var request model.CreateCommentRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userId := c.GetInt64("userId")
	postIdStr := c.Param("postId")
	postId, err := strconv.ParseInt(postIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = service.CommentService.CreateComment(postId, userId, &request)
	if err != nil {
		status := apperrors.ToHTTPStatus(err)
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Comment created successfully!"})
}
