package permission

import (
	"alpha/config"
	"alpha/handler"
	"alpha/pkg/errno"
	"alpha/repositories/data-mappers/model"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Create(c *gin.Context) {
	//声明 CreateRequest类型的变量
	var r CreateRequest

	//url获取并赋值
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
	//赋值
	p := model.PermissionModel{
		Label:     r.Label,
		Pid:       r.Pid,
		Level:     r.Level,
		Url:       r.Url,
		Sort:      r.Sort,
		Cond:      r.Cond,
		Component: r.Component,
		Icon:      r.Icon,
	}

	//创建
	if err := p.Create(); err != nil {
		config.Logger.Error("permission create",
			zap.Error(err),
		)
		handler.SendBadResponse(c, err, nil)
		return
	}
	//返回定义好的消息

	handler.SendResponse(c, nil, nil)

}
