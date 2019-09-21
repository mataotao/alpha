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

func UpdatePwd(c *gin.Context) {
	var r PwdRequest
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
		handler.SendBadResponseErrors(c, err, nil, errMap)
		return
	}
	user := &model.UserModel{
		BaseModel: model.BaseModel{
			Id: r.Id,
		},
		Password: r.Password,
	}
	if err := service.UpdatePwd(user); err != nil {
		config.Logger.Error("user update pwd",
			zap.Error(err),
		)
		handler.SendBadResponse(c, err, nil)
		return
	}

	handler.SendResponse(c, nil, nil)
}
