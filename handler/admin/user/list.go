package user

import (
	"alpha/config"
	"alpha/handler"
	"alpha/repositories/data-mappers/model"
	service "alpha/services/admin/user"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func List(c *gin.Context) {
	var r ListRequest
	if err := c.ShouldBindQuery(&r); err != nil {
		handler.SendBadResponse(c, err, nil)
		return
	}
	if _, err := govalidator.ValidateStruct(&r); err != nil {
		errMap := govalidator.ErrorsByField(err)
		handler.SendBadResponseErrors(c, err, nil, errMap)
		return
	}
	user := &model.UserModel{
		Username: r.Username,
		Name:     r.Name,
		Mobile:   r.Mobile,
	}
	info, err := service.List(user, r.Page, r.Limit)
	if err != nil {
		config.Logger.Error("user list",
			zap.Error(err),
		)
		handler.SendBadResponse(c, err, nil)
		return
	}

	handler.SendResponse(c, nil, info)
}
