package permission

import (
	"alpha/config"
	"alpha/handler"
	service "alpha/services/admin/permission"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func List(c *gin.Context) {
	userId := uint64(c.GetInt("user_id"))
	infos, err := service.List(userId)
	if err != nil {
		config.Logger.Error("permission list",
			zap.Error(err),
		)
		handler.SendBadResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, infos)
}
