package user

import (
	"alpha/config"
	userDomain "alpha/domain/entity/admin/user"
	"alpha/repositories/data-mappers/model"

	"go.uber.org/zap"
)

func UpdatePwd(user *model.UserModel) error {
	userEntity := userDomain.NewEntity(user.Id)
	userEntity.UserModel = *user
	//加密密码
	if err := userEntity.Encrypt(); err != nil {
		config.Logger.Error("user create",
			zap.Error(err),
		)
		return err

	}
	//更新信息
	if err := userEntity.UpdatePwd(); err != nil {
		config.Logger.Error("user update",
			zap.Error(err),
		)
		return err
	}
	return nil
}
