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

	"time"
)

func Create(c *gin.Context) {
	var r CreateRequest
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
		Username: r.Username,
		Name:     r.Name,
		Mobile:   r.Mobile,
		Password: r.Password,
		HeadImg:  r.HeadImg,
		LastTime: time.Now(),
		LastIp:   c.ClientIP(),
		Status:   model.ON,
	}
	if err := service.Create(user, r.RoleIds); err != nil {
		config.Logger.Error("role create",
			zap.Error(err),
		)
		handler.SendBadResponse(c, err, nil)
		return
	}

	handler.SendResponse(c, nil, nil)
}
