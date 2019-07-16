package login

import (
	"alpha/config"
	userDomain "alpha/domain/entity/admin/user"
	"alpha/pkg/errno"
	userService "alpha/services/admin/user"

	"go.uber.org/zap"

	"fmt"
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
	permissionIds, err := userService.PermissionIds(userEntity)
	if err != nil {
		config.Logger.Error("login in",
			zap.Error(err),
		)
		return false, err
	}

	fmt.Println(permissionIds)
	return false, nil
}
