package user

import (
	"alpha/config"
	userDomain "alpha/domain/entity/admin/user"
	"alpha/repositories/data-mappers/model"

	"go.uber.org/zap"
)

func Update(user *model.UserModel, roleIds []uint64) error {
	userEntity := userDomain.NewEntity(user.Id)
	userEntity.UserModel = *user
	if err := userEntity.Update(roleIds); err != nil {
		config.Logger.Error("user update",
			zap.Error(err),
		)
		return err
	}
	return nil
}
