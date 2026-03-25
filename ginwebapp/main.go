package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// secrets: registered user db
var secrets = gin.H{
	"foo":  gin.H{"email": "foo@bar.com", "phone": "123456898"},
	"subu": gin.H{"email": "subu@bar.com", "phone": "565656"},
	"lena": gin.H{"email": "lena@bar.com", "phone": "3454545"},
}

func main() {
	router := gin.Default()
	authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":  "bar",
		"subu": "1234",
		"lena": "lena",
	}))
	//end points - /admin/secrets
	authorized.GET("/secrets", func(c *gin.Context) {
		//get user, it was set by the BasicAuth Middleware
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "No secret"})
		}
	})

	router.Run(":8080")
}
