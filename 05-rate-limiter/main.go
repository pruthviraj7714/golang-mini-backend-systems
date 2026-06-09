package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Request struct {
	requestCount    int
	lastRequestTime int64
}

type RateLimiter struct {
	requests map[string]Request
}

var rateLimiter = RateLimiter{
	requests: make(map[string]Request),
}

func middleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		ip := c.ClientIP()
		req := rateLimiter.requests[ip]
		now := time.Now().Unix()

		if now >= req.lastRequestTime+60 {
			req.requestCount = 0
			req.lastRequestTime = now
		} else {
			req.requestCount++

			if req.requestCount > 10 {
				c.JSON(429, gin.H{
					"message": "Too many requests",
				})
				c.Abort()
				return
			}
		}

		rateLimiter.requests[ip] = req
		c.Next()
	}
}

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	authRouter := r.Group("/auth")
	{
		authRouter.Use(middleware())
		authRouter.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	r.Run()
}
