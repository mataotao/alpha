package middleware

import (
	"alpha/handler"
	"alpha/pkg/errno"
	"alpha/pkg/token"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		if _, err := token.ParseRequest(c); err != nil {
			handler.SendResponse(c, errno.InternalServerError, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
