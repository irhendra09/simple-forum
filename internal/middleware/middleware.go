package middleware

import (
	"net/http"

	"donedev.com/simple-forum/internal/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "No Authorization header",
			})
			c.Abort()
			return
		}

		userId, err := utils.ParseToken(authHeader)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		c.Set("userId", userId)
		c.Next()
	}
}
