package main

import (
	"donedev.com/simple-forum/internal/configs"
	"donedev.com/simple-forum/internal/model"
	"donedev.com/simple-forum/internal/repository"
	"donedev.com/simple-forum/internal/routes"
	"donedev.com/simple-forum/internal/service"
	"donedev.com/simple-forum/internal/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	// try to load optional YAML configuration first
	_ = configs.LoadConfig("config.yaml")

	db := configs.ConnectDB()

	// instantiate repositories
	userRepo := repository.NewGormUserRepository(db)
	postRepo := repository.NewGormPostRepository(db)
	commentRepo := repository.NewGormCommentRepository(db)

	// token service
	tokenSvc := utils.NewJwtTokenService([]byte(configs.GetJWTSecret()))
	utils.TokenSvc = tokenSvc

	// instantiate refresh token repository
	refreshRepo := repository.NewGormRefreshRepository(db)

	// auto-migrate refresh token table (creates table if missing)
	if err := db.AutoMigrate(&model.RefreshToken{}); err != nil {
		panic(err)
	}

	// instantiate services and assign package vars
	service.NewUserService(userRepo, tokenSvc, refreshRepo)
	service.NewPostService(postRepo)
	service.NewCommentService(commentRepo)

	r := gin.Default()
	routes.UserRoutes(r)
	routes.PostRoutes(r)
	r.Run(":8080")
}
