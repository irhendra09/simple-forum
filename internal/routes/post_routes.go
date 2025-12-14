package routes

import (
	"donedev.com/simple-forum/internal/handler"
	"donedev.com/simple-forum/internal/middleware"
	"github.com/gin-gonic/gin"
)

func PostRoutes(r *gin.Engine) {
	router := r.Group("/post")
	router.Use(middleware.AuthMiddleware())
	router.POST("/", handler.CreatePostsHandler)
	router.GET("/", handler.GetAllPostsHandler)
	router.POST("comment/:postId", handler.CreateCommentHandler)
	router.POST("user-activity/:postId", handler.UserActivity)
	router.GET("/:postId", handler.GetPostByIdHandler)
}
