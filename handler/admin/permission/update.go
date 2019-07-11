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
	//赋值
	p := model.PermissionModel{
		Label:         r.Label,
		IsContainMenu: r.IsContainMenu,
		Url:           r.Url,
		Sort:          r.Sort,
		Cond:          r.Cond,
		Icon:          r.Icon,
	}
	p.Id = r.Id //创建
	if err := p.Update(); err != nil {
		config.Logger.Error("permission update",
			zap.Error(err),
		)
		handler.SendBadResponse(c, err, nil)
		return
	}
	//返回定义好的消息

	handler.SendResponse(c, nil, nil)

}
