package limiter

import (
	log "alpha/config"
	redis "alpha/repositories/data-mappers/go-redis"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	limiter "github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	sredis "github.com/ulule/limiter/v3/drivers/store/redis"
	"go.uber.org/zap"
)

// * 5 reqs/second: "5-S"
// * 10 reqs/minute: "10-M"
// * 1000 reqs/hour: "1000-H"
// * 2000 reqs/day: "2000-D"

func UIP() gin.HandlerFunc {
	server := viper.GetString("name")
	rate, err := limiter.NewRateFromFormatted(viper.GetString("limiter.rate"))
	if err != nil {
		log.Logger.Error("limiter init",
			zap.Error(err),
		)
		return nil
	}

	store, err := sredis.NewStoreWithOptions(redis.Client.Client, limiter.StoreOptions{
		Prefix:   fmt.Sprintf("limiter_%s", server),
		MaxRetry: 3,
	})
	if err != nil {
		log.Logger.Error("limiter init",
			zap.Error(err),
		)
		return nil
	}
	middleware := mgin.NewMiddleware(limiter.New(store, rate))
	return middleware
}
