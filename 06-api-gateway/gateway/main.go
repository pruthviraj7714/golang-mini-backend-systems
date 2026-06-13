package main

import (
	"api-gateway/middleware"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	target1, err := url.Parse("http://localhost:8081")

	if err != nil {
		log.Fatalf("failed to parse URL: %v", err)
	}

	target2, err := url.Parse("http://localhost:8082")

	if err != nil {
		log.Fatalf("failed to parse URL: %v", err)
	}

	userProxy := httputil.NewSingleHostReverseProxy(target1)
	productProxy := httputil.NewSingleHostReverseProxy(target2)

	r.Use(middleware.LoggingMiddleware())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "gateway is running",
		})
	})

	r.Any("/api/users/*path", func(c *gin.Context) {
		path := c.Param("path")

		if path == "/" {
			c.Request.URL.Path = "/users"
		} else {
			c.Request.URL.Path = "/users" + path
		}

		userProxy.ServeHTTP(c.Writer, c.Request)
	})

	r.GET("/api/products", func(c *gin.Context) {
		path := c.Param("path")

		if path == "/" {
			c.Request.URL.Path = "/products"
		} else {
			c.Request.URL.Path = "/products" + path
		}

		productProxy.ServeHTTP(c.Writer, c.Request)
	})

	r.Run(":8080")
}
