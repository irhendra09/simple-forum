package main

import (
	"donedev.com/simple-forum/internal/configs"
	"donedev.com/simple-forum/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	configs.ConnectDB()
	r := gin.Default()
	routes.UserRoutes(r)
	routes.PostRoutes(r)
	r.Run(":8080")
}
