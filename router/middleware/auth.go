package middleware

import (
	"alpha/pkg/errno"
	"alpha/pkg/token"
	"fmt"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		t, err := token.ParseRequest(c)
		if err != nil {
			c.String(401, "")
			c.Error(errno.ErrTokenInvalid)
			c.Abort()
			return
		}
		c.Set("user_id", t.ID)
		c.Set("user_id", t.Username)
		fmt.Println(c.HandlerName())
		c.Next()
	}
}
