package routes

import (
	"donedev.com/simple-forum/internal/handler"
	"donedev.com/simple-forum/internal/middleware"
	"github.com/gin-gonic/gin"
)

func PostRoutes(r *gin.Engine) {
	r.Use(middleware.AuthMiddleware())
	r.POST("/post", handler.CreatePostsHandler)
}
