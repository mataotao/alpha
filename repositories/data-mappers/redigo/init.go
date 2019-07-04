package redigo

import (
	log "alpha/config"
	"bytes"
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"strconv"
)

type InitPool struct {
	Pool *redis.Pool
}

var Pool *InitPool

func GetPool() *redis.Pool {
	var redisConnect bytes.Buffer
	redisConnect.WriteString(viper.GetString("redis.host"))
	redisConnect.WriteString(":")
	redisConnect.WriteString(strconv.Itoa(viper.GetInt("redis.port")))
	db := viper.GetInt("redis.db")
	pwd := viper.GetString("redis.pwd")
	pool := &redis.Pool{
		MaxIdle:     viper.GetInt("redis.maxidle"),
		MaxActive:   viper.GetInt("redis.maxactive"), //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: 300,                             //连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) { //要连接的redis数据库
			c, err := redis.Dial("tcp", redisConnect.String())
			if err != nil {
				log.Logger.Error("redis-init",
					zap.Error(err),
				)
				return nil, err
			}
			if _, err := c.Do("AUTH", pwd); err != nil {
				log.Logger.Error("redis-init",
					zap.Error(err),
				)
				c.Close()
				return nil, err
			}
			if _, err := c.Do("SELECT", db); err != nil {
				log.Logger.Error("redis-init",
					zap.Error(err),
				)
				c.Close()
				return nil, err
			}
			return c, err
		},
	}
	return pool
}
func (pool *InitPool) Init() {
	Pool = &InitPool{
		Pool: GetPool(),
	}
}

func (pool *InitPool) Close() {
	Pool.Pool.Close()
}
