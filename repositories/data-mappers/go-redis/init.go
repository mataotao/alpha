package redigo

import (
	log "alpha/config"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Init struct {
	Client *redis.Client
}

var Client *Init

func NewClient() *redis.Client {
	options := &redis.Options{
		Addr:     viper.GetString("redis.addr"), // use default Addr
		Password: viper.GetString("redis.pwd"),  // no password set
		DB:       viper.GetInt("redis.db"),      // use default DB
	}
	client := redis.NewClient(options)
	_, err := client.Ping().Result()
	if err != nil {
		log.Logger.Error("redis ping error",
			zap.Error(err),
		)
	}

	return client
}
func (pool *Init) Init() {
	Client = &Init{
		Client: NewClient(),
	}
}

func (pool *Init) Close() {
	Client.Client.Close()
}
