package permission

import (
	"alpha/config"
	"alpha/handler"
	service "alpha/services/admin/permission"

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
		handler.SendBadResponseErrors(c, err, nil, errMap)
		return
	}
	info, err := service.Get(r.Id)
	if err != nil {
		config.Logger.Error("permission get",
			zap.Error(err),
		)
		//错误返回
		handler.SendBadResponse(c, err, nil)
		return
	}
	//返回权限数据
	handler.SendResponse(c, nil, info)

}
