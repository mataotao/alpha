package login

import (
	"alpha/config"
	userDomain "alpha/domain/entity/admin/user"
	"alpha/pkg/auth"
	"alpha/pkg/errno"

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
	if err := auth.Compare(userEntity.UserModel.Password, pwd); err != nil {
		config.Logger.Error("login in",
			zap.Error(err),
		)
		return false, errno.ErrUserNameOrPassword
	}
	if userEntity.IsFreeze() {
		return false, errno.UserFreeze
	}

	fmt.Println(userEntity)
	return false, nil
}
