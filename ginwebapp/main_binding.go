package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

//	type Login struct {
//		Username string `json:"username" binding:"required"`
//		Password string `json:"password" binding:"required"`
//	}
type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=8,max=10"`
}

func loginHandler(c *gin.Context) {
	var body Login
	//MustBindJson
	//1.parase the json body into the body struct
	//if it fails (eg missing fields or invalid json),it sets error message
	if err := c.BindJSON(&body); err != nil {
		return
	}
	//send response
	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"status": "success",
		"data":   body.Username,
	})

}
func getFrindlyError(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "max":
		return "The maximum length field must be greater than zero"
	case "min":
		return "The minimum length field must be greater than zero"
	case "email":
		return "The email address is required"
	case "alphanum":
	case "required-space":
		return "The filed cant be alphanumeric"
	}
	return "Invalid value"
}
func LoginHandlerWithError(c *gin.Context) {
	var body Login
	if err := c.ShouldBind(&body); err != nil {
		var ve validator.ValidationErrors
		//Check if the error is actually a validation error
		if errors.As(err, &ve) {
			out := make(map[string]string)
			for _, fe := range ve {
				//we map the json field name to our friendly error message
				out[fe.Field()] = getFrindlyError(fe)
			}
			c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": out})
			return
		}
		//Handle non validation errors like Malformed json
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json format"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "logged In!"})

}

func main() {
	router := gin.Default()
	router.POST("/login", loginHandler)
	router.POST("/logincustom", LoginHandlerWithError)
	fmt.Println("Server is running on port 8080")
	router.Run(":8080")
}
