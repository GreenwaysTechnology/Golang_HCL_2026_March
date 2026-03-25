package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/welcome", func(c *gin.Context) {
		name := c.DefaultQuery("firstname", "Guest")
		c.String(http.StatusOK, "Hello %s", name)
	})
	fmt.Println("Server is running on port 8080")
	router.Run(":8080")
}
