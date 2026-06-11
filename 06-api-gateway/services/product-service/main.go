package productservice

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/product", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"service": "product-service",
		})
	})

	fmt.Println("Product Service is running on Port 8081")
	r.Run(":8082")
}
