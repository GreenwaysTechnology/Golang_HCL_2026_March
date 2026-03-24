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
			"message": "Home",
		})
	})
	router.GET("/users", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Users GET",
		})
	})
	router.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"message": "Users GET By Param" + id,
		})
	})
	router.POST("/users", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Users POST",
		})
	})
	router.PUT("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"message": "Users PUT By Param" + id,
		})
	})
	router.DELETE("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"message": "Users DELETE By Param" + id,
		})
	})

	fmt.Println("Gin Server Started...")
	router.Run()
}
