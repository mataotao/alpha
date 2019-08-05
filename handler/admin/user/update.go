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

func Update(c *gin.Context) {
	var r UpdateRequest
	if err := c.ShouldBindUri(&r); err != nil {
		handler.SendBadResponse(c, err, nil)
		return
	}
	if err := c.ShouldBind(&r); err != nil {
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
		Name:   r.Name,
		Mobile: r.Mobile,
		Avatar: r.Avatar,
	}
	if err := service.Update(user, r.RoleIds); err != nil {
		config.Logger.Error("user update",
			zap.Error(err),
		)
		handler.SendBadResponse(c, err, nil)
		return
	}

	handler.SendResponse(c, nil, nil)
}
