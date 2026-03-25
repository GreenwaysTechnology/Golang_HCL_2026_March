package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// define struct
type UserRequest struct {
	Username string `json:"username" binding:"required,is-admin"`
	Age      int    `json:"age" binding:"required,gte=18"`
}

// define custom validation logic
var validateGopher validator.Func = func(fl validator.FieldLevel) bool {
	//get field value
	value := fl.Field().String()
	return strings.Contains(strings.ToLower(value), "admin")
}

func main() {
	r := gin.Default()
	//Register custom validator with Gin engine
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("is-admin", validateGopher)
	}
	r.POST("/register", func(c *gin.Context) {
		var user UserRequest
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to," + user.Username})
	})

	fmt.Println("Server started at port 8080")
	r.Run(":8080")
}
