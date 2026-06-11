package userservice

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"service": "user-service",
		})
	})

	fmt.Println("User Service is running on Port 8081")
	r.Run(":8081")
}
