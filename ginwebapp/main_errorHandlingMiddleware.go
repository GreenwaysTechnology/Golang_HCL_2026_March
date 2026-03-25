package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Process the request first
		c.Next()
		//Check if any errors were added to the context
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": err.Error(),
			})
		}
	}
}
func main() {
	router := gin.Default()
	//attach error handler global middleware
	router.Use(ErrorHandler())
	router.GET("/ok", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	router.GET("/error", func(c *gin.Context) {
		//simulate error
		c.Error(errors.New("Something went wrong"))
	})
	router.Run(":8080")
}
