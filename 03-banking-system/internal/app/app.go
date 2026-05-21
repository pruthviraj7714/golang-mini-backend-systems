package app

import (
	"banking-system/internal/config"
	"banking-system/internal/db"
	"banking-system/internal/handlers"
	"banking-system/internal/repository"
	"banking-system/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Start() {

	cfg := config.LoadConfig()

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	})

	database := db.Connect(cfg.DBURL)

	userRepo := &repository.UserRepository{
		DB: database,
	}

	jwtService := services.NewJWTService(
		cfg.AccessTokenSecret,
		cfg.RefreshTokenSecret,
	)

	userService := &services.UserService{
		Repo:       userRepo,
		JWTService: jwtService,
	}

	userHandler := &handlers.UserHandler{
		UserService: userService,
	}

	authRouter := router.Group("/auth")
	authRouter.POST("/register", userHandler.Register)
	authRouter.POST("/login", userHandler.Login)
	router.Run(":" + cfg.Port)

}
