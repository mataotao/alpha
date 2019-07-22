package middleware

import (
	permissionDomain "alpha/domain/entity/admin/permission"
	userDomain "alpha/domain/entity/admin/user"
	"alpha/pkg/errno"
	"alpha/pkg/token"
	redis "alpha/repositories/data-mappers/go-redis"
	gredis "github.com/go-redis/redis"
	"strconv"

	"fmt"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		t, err := token.ParseRequest(c)
		if err != nil {
			c.String(401, errno.ErrTokenInvalid.Message)
			c.Error(errno.ErrTokenInvalid)
			c.Abort()
			return
		}
		c.Set("user_id", t.ID)
		c.Set("user_id", t.Username)
		//获取权限id
		value, err := redis.Client.Client.Get(fmt.Sprintf(permissionDomain.CachePermissionKey, c.HandlerName())).Result()
		if err == gredis.Nil {
			c.String(401, errno.ErrAuthInvalid.Message)
			c.Error(errno.ErrTokenInvalid)
			c.Abort()
			return
		}
		id, err := strconv.Atoi(value)
		if err != nil {
			c.String(401, errno.ErrAuthInvalid.Message)
			c.Error(errno.ErrTokenInvalid)
			c.Abort()
			return
		}
		auth, err := redis.Client.Client.GetBit(fmt.Sprintf(userDomain.PermissionKey, t.ID), int64(id)).Result()
		if err != nil {
			c.String(401, errno.ErrAuthInvalid.Message)
			c.Error(errno.ErrTokenInvalid)
			c.Abort()
			return
		}
		if auth == 0 {
			c.String(401, errno.ErrAuthInvalid.Message)
			c.Error(errno.ErrTokenInvalid)
			c.Abort()
			return
		}
		c.Next()
	}
}
