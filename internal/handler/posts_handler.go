package handler

import (
	"net/http"
	"strconv"

	apperrors "donedev.com/simple-forum/internal/errors"
	"donedev.com/simple-forum/internal/model"
	"donedev.com/simple-forum/internal/service"
	"github.com/gin-gonic/gin"
)

func CreatePostsHandler(c *gin.Context) {
	var request model.CreatePostRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userId := c.GetInt64("userId")
	err := service.PostService.CreatePost(&request, userId)
	if err != nil {
		status := apperrors.ToHTTPStatus(err)
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "successfully created post",
	})
}

func UserActivity(c *gin.Context) {
	var request model.UserActivityRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	postIdStr := c.Param("postId")
	postId, err := strconv.ParseInt(postIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userId := c.GetInt64("userId")
	err = service.PostService.UpsertUserActivity(request, postId, userId)
	if err != nil {
		status := apperrors.ToHTTPStatus(err)
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func GetAllPostsHandler(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	sizeStr := c.DefaultQuery("size", "10")

	page, _ := strconv.Atoi(pageStr)
	size, _ := strconv.Atoi(sizeStr)

	result, err := service.PostService.GetAllPosts(page, size)
	if err != nil {
		status := apperrors.ToHTTPStatus(err)
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}
	c.JSON(
		http.StatusOK,
		result,
	)
}

func GetPostByIdHandler(c *gin.Context) {
	postIdStr := c.Param("postId")
	postId, err := strconv.ParseInt(postIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": apperrors.ErrBadRequest.Error()})
		return
	}

	response, err := service.PostService.GetPostById(postId)
	if err != nil {
		status := apperrors.ToHTTPStatus(err)
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)

}
