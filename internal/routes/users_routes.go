package routes

import (
	"donedev.com/simple-forum/internal/handler"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	router := r.Group("/users")
	router.POST("/signup", handler.SignUp)
	router.POST("/login", handler.SignIn)
	router.POST("/refresh", handler.RefreshToken)
	router.POST("/logout", handler.Logout)
}
