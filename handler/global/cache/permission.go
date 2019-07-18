package cache

import (
	"alpha/config"
	"alpha/handler"
	service "alpha/services/global/cache"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Permission(c *gin.Context) {
	err := service.Permission(c.GetString("X-Request-Id"))
	if err != nil {
		config.Logger.Error("cache permission",
			zap.Error(err),
		)
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
}
