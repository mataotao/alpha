package permission

import (
	"alpha/config"
	userDomain "alpha/domain/entity/admin/user"
	"alpha/pkg/errno"
	"alpha/repositories/data-mappers/model"
	userService "alpha/services/admin/user"

	"go.uber.org/zap"
)

func List(userId uint64) ([]*model.PermissionListInfo, error) {
	list := make([]*model.PermissionListInfo, 0)
	//用户权限过滤
	userEntity := userDomain.NewEntity(userId)
	notfound, err := userEntity.Get("*")
	if err != nil {
		config.Logger.Error("permission list",
			zap.Error(err),
		)
		return list, err
	}
	if notfound {
		return list, errno.ErrDBNotFoundRecord
	}
	//获取权限id
	if _, err := userService.PermissionIds(userEntity); err != nil {
		config.Logger.Error("login in",
			zap.Error(err),
		)
		return list, err
	}
	//获取权限列表
	p := new(model.PermissionModel)
	plist, err := p.AllByIds("*", userEntity.PermissionIds)
	if err != nil {
		config.Logger.Error("permission get",
			zap.Error(err),
		)
		return list, err
	}
	list = p.RecursivePermission(0, plist)
	return list, nil
}
