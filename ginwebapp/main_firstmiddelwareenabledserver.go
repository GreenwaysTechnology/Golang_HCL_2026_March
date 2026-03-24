package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	//initalize the gin Engine with middleware support
	router := gin.Default()
	//router with home page
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	//router.Run(":8080")
	//default port is 8080
	fmt.Println("Gin Server Started...")
	router.Run()

}
