package user

import (
	"alpha/config"
	userDomain "alpha/domain/entity/admin/user"
	"alpha/pkg/errno"
	"alpha/repositories/data-mappers/model"

	"go.uber.org/zap"
)

func Create(user *model.UserModel, roleIds []uint64) error {
	userEntity := userDomain.NewEntity(0)
	userEntity.UserModel = *user
	//检查用户名唯一
	if userEntity.UserModel.Unique() == false {
		return errno.ErrUserNameNotUnique
	}
	//加密密码
	if err := userEntity.UserModel.Encrypt(); err != nil {
		config.Logger.Error("user create",
			zap.Error(err),
		)
		return err

	}
	if err := userEntity.UserModel.Create(roleIds); err != nil {
		config.Logger.Error("user create",
			zap.Error(err),
		)
		return err
	}
	return nil
}
