package role

import (
	"alpha/config"
	"alpha/handler"
	"alpha/pkg/errno"
	"alpha/repositories/data-mappers/model"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Update(c *gin.Context) {
	var r UpdateRequest

	//绑定id  uri方式
	if err := c.ShouldBindUri(&r); err != nil {
		//返回错误
		handler.SendBadResponse(c, err, nil)
		return
	}
	//绑定提交的值
	if err := c.ShouldBind(&r); err != nil {
		//返回错误
		handler.SendBadResponse(c, err, nil)
		return
	}
	//验证
	if _, err := govalidator.ValidateStruct(&r); err != nil {
		errMap := govalidator.ErrorsByField(err)
		handler.SendBadResponseErrors(c, errno.ErrValidation, nil, errMap)
		return
	}
	rm := &model.RoleModel{
		Name:        r.Name,
		Description: r.Description,
	}
	rm.Id = r.Id
	if err := rm.Update(r.Permission); err != nil {
		config.Logger.Error("role update",
			zap.Error(err),
		)
		handler.SendBadResponse(c, nil, nil)
		return
	}
	handler.SendResponse(c, nil, nil)

}
