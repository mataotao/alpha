package permission

import (
	"alpha/config"
	"alpha/handler"
	"alpha/repositories/data-mappers/model"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Delete(c *gin.Context) {
	var r DeleteRequest
	//绑定id  uri方式
	if err := c.ShouldBindUri(&r); err != nil {
		handler.SendBadResponse(c, err, nil)
		return
	}
	//验证
	if _, err := govalidator.ValidateStruct(&r); err != nil {
		errMap := govalidator.ErrorsByField(err)
		handler.SendBadResponseErrors(c, err, nil, errMap)
		return
	}
	p := new(model.PermissionModel)
	p.Id = r.Id
	if err := p.Delete(); err != nil {
		config.Logger.Error("permission delete",
			zap.Error(err),
		)
		//返回错误
		handler.SendBadResponse(c, err, nil)
		return
	}
	//返回成功
	handler.SendResponse(c, nil, nil)
}
