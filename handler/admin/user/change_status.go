package user

import (
	"alpha/config"
	"alpha/handler"
	"alpha/pkg/errno"
	"alpha/repositories/data-mappers/model"
	service "alpha/services/admin/user"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ChangeStatus(c *gin.Context) {
	var r StatusRequest
	if err := c.ShouldBindUri(&r); err != nil {
		handler.SendBadResponse(c, err, nil)
		return
	}
	if _, err := govalidator.ValidateStruct(&r); err != nil {
		errMap := govalidator.ErrorsByField(err)
		handler.SendBadResponseErrors(c, errno.ErrValidation, nil, errMap)
		return
	}
	user := &model.UserModel{
		BaseModel: model.BaseModel{
			Id: r.Id,
		},
	}
	if err := service.ChangeStatus(user); err != nil {
		config.Logger.Error("role change status",
			zap.Error(err),
		)
		handler.SendBadResponse(c, err, nil)
		return
	}

	handler.SendResponse(c, nil, nil)
}
