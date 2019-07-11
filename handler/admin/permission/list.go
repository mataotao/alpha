package permission

import (
	"alpha/config"
	"alpha/handler"
	service "alpha/services/permission"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func List(c *gin.Context) {
	infos, err := service.List()
	if err != nil {
		config.Logger.Error("permission list",
			zap.Error(err),
		)
		handler.SendBadResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, infos)
}
