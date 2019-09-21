package role

import (
	"alpha/config"
	"alpha/handler"
	"alpha/repositories/data-mappers/model"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Create(c *gin.Context) {
	var r CreateRequest
	if err := c.ShouldBind(&r); err != nil {
		handler.SendBadResponse(c, err, nil)
		return
	}
	if _, err := govalidator.ValidateStruct(&r); err != nil {
		errMap := govalidator.ErrorsByField(err)
		handler.SendBadResponseErrors(c, err, nil, errMap)
		return
	}
	role := model.RoleModel{
		Name:        r.Name,
		Description: r.Description,
	}
	if err := role.Create(r.Permission); err != nil {
		config.Logger.Error("role create",
			zap.Error(err),
		)
		handler.SendBadResponse(c, err, nil)
		return
	}

	handler.SendResponse(c, nil, nil)
}
