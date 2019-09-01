package limiter

import (
	"alpha/config"
	redis "alpha/repositories/data-mappers/go-redis"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"fmt"
)

const AdminRedisCellKey = "alpha:redis:cell:admin:%d"

func AdminRedisCell() gin.HandlerFunc {
	return func(c *gin.Context) {
		k := fmt.Sprintf(AdminRedisCellKey, c.GetInt("user_id"))
		//CL.THROTTLE <key> <max_burst> <count per period> <period> [<quantity>]
		res, err := redis.Client.Client.Do("CL.THROTTLE", k, 15, 30, 60).Result()
		if err != nil {
			config.Logger.Error("admin redis-cell",
				zap.Error(err),
			)
			c.String(429, "You have reached maximum request limit.")
			c.Abort()
			return
		}
		if res.([]interface{})[0].(int64) == 1 {
			c.String(429, "You have reached maximum request limit.")
			c.Abort()
			return
		}
		c.Next()
	}
}
