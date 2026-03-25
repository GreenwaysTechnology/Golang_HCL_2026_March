package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// route specific middleware - pingpongMiddeware
func PingPongMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Next() //very important
		latency := time.Since(t)
		log.Printf("PingPong Latency :%s", latency)
	}
}

// auth middleware
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		fmt.Println(token)
		if token != "secret-token" {
			fmt.Println("Auth Middleware Failed")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		fmt.Println("Auth Middleware Success")
		//Continue to the Next handler
		c.Next()
	}
}

func main() {
	router := gin.New()
	//Global middlewares
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	//Route specific middleware
	//router.method("url","middleare",handler)
	router.GET("/ping", PingPongMiddleware(), func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	//group middlware : for set of urls
	//create group
	authorized := router.Group("/")
	//attach middleware for autorized group
	authorized.Use(AuthMiddleware())
	{
		authorized.POST("/login", func(c *gin.Context) {
			c.String(http.StatusOK, "login")
		})
		authorized.POST("/logout", func(c *gin.Context) {
			c.String(http.StatusOK, "logout")
		})
		authorized.POST("/refresh", func(c *gin.Context) {
			c.String(http.StatusOK, "refresh")
		})
		authorized.POST("/refresh-token", func(c *gin.Context) {
			c.String(http.StatusOK, "refresh-token")
		})
		authorized.POST("/register", func(c *gin.Context) {
			c.String(http.StatusOK, "register")
		})
	}
	router.Run(":8080")
}
