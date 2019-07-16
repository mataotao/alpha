package login

import (
	"alpha/config"
	userDomain "alpha/domain/entity/admin/user"
	"alpha/pkg/errno"
	userService "alpha/services/admin/user"

	"go.uber.org/zap"
)

func In(username, pwd, ip string) (bool, error) {
	userEntity := userDomain.NewEntityS(username)
	notfound, err := userEntity.Get()
	if err != nil {
		config.Logger.Error("login in",
			zap.Error(err),
		)
		return false, err
	}
	if notfound == true {
		return false, errno.ErrDBNotFoundRecord
	}
	//检测密码
	if err := userEntity.Compare(pwd); err != nil {
		config.Logger.Error("login in",
			zap.Error(err),
		)
		return false, errno.ErrUserNameOrPassword
	}
	//检测用户冻结
	if userEntity.IsFreeze() {
		return false, errno.ErrUserFreeze
	}
	//获取权限id
	if _, err := userService.PermissionIds(userEntity); err != nil {
		config.Logger.Error("login in",
			zap.Error(err),
		)
		return false, err
	}
	//权限id写入redis 位图
	if err := userEntity.SetPermissionToCache(); err != nil {
		config.Logger.Error("login in",
			zap.Error(err),
		)
		return false, err
	}

	return false, nil
}
