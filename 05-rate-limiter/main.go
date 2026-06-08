package main

import (
	"fmt"
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

		fmt.Println("ip: " + ip)

		if _, exists := rateLimiter.requests[ip]; !exists {
			rateLimiter.requests[ip] = Request{
				requestCount:    1,
				lastRequestTime: time.Now().Unix(),
			}
		} else {
			currentRequest := rateLimiter.requests[ip]

			if currentRequest.requestCount >= 10 && time.Now().Unix() <= currentRequest.lastRequestTime+60 {
				c.JSON(429, gin.H{
					"message": "Too many requests",
				})
				return
			}

			rateLimiter.requests[ip] = Request{
				requestCount:    rateLimiter.requests[ip].requestCount + 1,
				lastRequestTime: time.Now().Unix(),
			}
		}

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

	r.GET("/ping", middleware(), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
