package user

import (
	"alpha/config"
	"alpha/handler"
	service "alpha/services/admin/user"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Information(c *gin.Context) {
	userId := uint64(c.GetInt("user_id"))
	info, err := service.Get(userId)
	if err != nil {
		config.Logger.Error("user get",
			zap.Error(err),
		)
		handler.SendBadResponse(c, err, nil)
		return
	}

	handler.SendResponse(c, nil, info)
}
