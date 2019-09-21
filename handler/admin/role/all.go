package role

import (
	"alpha/config"
	"alpha/handler"
	service "alpha/services/admin/role"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func All(c *gin.Context) {

	list, err := service.All()
	if err != nil {
		config.Logger.Error("role all",
			zap.Error(err),
		)
		handler.SendBadResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, list)

}
