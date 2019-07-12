package permission

import (
	"alpha/config"
	"alpha/repositories/data-mappers/model"
	"go.uber.org/zap"
)

func List() ([]*model.PermissionListInfo, error) {
	list := make([]*model.PermissionListInfo, 0)
	//TODO 用户权限过滤
	//获取权限列表
	p := new(model.PermissionModel)
	plist, err := p.AllByIds("*", []uint64{})
	if err != nil {
		config.Logger.Error("permission get",
			zap.Error(err),
		)
		return list, err
	}
	list = p.RecursivePermission(0, plist)
	return list, nil
}
