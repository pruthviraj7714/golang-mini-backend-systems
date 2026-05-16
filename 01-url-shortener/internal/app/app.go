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

	cfg := config.LoadConfig()

	database := db.Connect()

	userRepository := &repository.UserRepository{
		DB: database,
	}

	userService := &services.UserService{
		UserRepository: userRepository,
	}

	userHandler := &handler.UserHandler{
		UserService: userService,
	}

	urlRepository := &repository.UrlRepository{
		DB: database,
	}

	urlService := &services.UrlService{
		UrlRepository: urlRepository,
	}

	urlHandler := &handler.UrlHandler{
		UrlService: urlService,
	}

	database.AutoMigrate(&models.User{})
	database.AutoMigrate(&models.Url{})

	authRouter := r.Group("/auth")

	authRouter.POST("/register", userHandler.RegisterUser)
	authRouter.POST("/login", userHandler.LoginUser)

	urlRouter := r.Group("/url")
	urlRouter.Use(middlewares.AuthMiddleware())

	urlRouter.POST("/shorten", urlHandler.CreateUrl)
	r.GET("/:shortId", urlHandler.RedirectToShortUrl)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})

	r.Run(":" + cfg.Port)
}
