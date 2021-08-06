package middleware

import "github.com/gin-gonic/gin"

func MyAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"root": "123456",
	})
}
