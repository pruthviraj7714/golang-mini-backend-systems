package app

import (
	"net/http"
	"url-shortener/internal/config"
	"url-shortener/internal/db"
	"url-shortener/internal/handler"
	"url-shortener/internal/middlewares"
	"url-shortener/internal/models"
	"url-shortener/internal/repository"
	"url-shortener/internal/services"

	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()

	database := db.Connect()

	userRepository := &repository.UserRepository{
		DB: database,
	}

	userServices := &services.UserServices{
		UserRepository: userRepository,
	}

	userHandler := &handler.UserHandler{
		UserServices: userServices,
	}

	database.AutoMigrate(&models.User{})

	authRouter := r.Group("/auth")

	authRouter.POST("/register", userHandler.RegisterUser)
	authRouter.POST("/login", userHandler.LoginUser)

	urlRouter := r.Group("/url")
	urlRouter.Use(middlewares.AuthMiddleware())

	urlRouter.POST("/shorten")

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})

	r.Run(":" + config.LoadConfig().Port)
}
