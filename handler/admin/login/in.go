package login

import (
	"alpha/config"
	"alpha/handler"
	"alpha/pkg/errno"
	service "alpha/services/admin/login"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func In(c *gin.Context) {
	var r InRequest
	if err := c.ShouldBind(&r); err != nil {
		//返回错误
		handler.SendBadResponse(c, err, nil)
		return
	}
	if _, err := govalidator.ValidateStruct(&r); err != nil {
		errMap := govalidator.ErrorsByField(err)
		handler.SendBadResponseErrors(c, errno.ErrValidation, nil, errMap)
		return
	}
	login, err := service.In(r.Username, r.Password, c.ClientIP())
	if err != nil {
		config.Logger.Error("login in",
			zap.Error(err),
		)
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, login)

}
