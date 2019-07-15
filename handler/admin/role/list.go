package role

import (
	"alpha/config"
	"alpha/handler"
	service "alpha/services/admin/role"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func List(c *gin.Context) {
	var r ListRequest
	//绑定id  uri方式
	if err := c.ShouldBindQuery(&r); err != nil {
		handler.SendBadResponse(c, err, nil)
		return
	}

	list, err := service.List(r.Name, r.Page, r.Limit)
	if err != nil {
		config.Logger.Error("role list",
			zap.Error(err),
		)
		handler.SendBadResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, list)

}
