package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	//initialize the gin Engine without middleware support
	router := gin.New() //router with home page
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
