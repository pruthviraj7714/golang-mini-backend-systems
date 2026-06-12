package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		fmt.Printf("%s %s", c.Request.Method, c.Request.URL)

	}
}
