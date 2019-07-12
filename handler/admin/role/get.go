package role

import (
	"alpha/config"
	"alpha/handler"
	"alpha/pkg/errno"
	service "alpha/services/admin/role"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Get(c *gin.Context) {
	var r GetRequest
	//绑定id  uri方式
	if err := c.ShouldBindUri(&r); err != nil {
		handler.SendBadResponse(c, err, nil)
		return
	}
	//验证
	if _, err := govalidator.ValidateStruct(&r); err != nil {
		errMap := govalidator.ErrorsByField(err)
		handler.SendBadResponseErrors(c, errno.ErrValidation, nil, errMap)
		return
	}

	info, err := service.Get(r.Id)
	if err != nil {
		config.Logger.Error("role get",
			zap.Error(err),
		)
		handler.SendBadResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, info)

}
