package user

import (
	"alpha/config"
	userDomain "alpha/domain/entity/admin/user"
	"alpha/repositories/data-mappers/model"

	"go.uber.org/zap"
)

func ChangeStatus(user *model.UserModel) error {
	userEntity := userDomain.NewEntity(user.Id)
	userEntity.UserModel = *user
	//更新信息
	if err := userEntity.ChangeStatus(); err != nil {
		config.Logger.Error("user change status",
			zap.Error(err),
		)
		return err
	}
	return nil
}
