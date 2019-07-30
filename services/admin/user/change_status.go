package user

import (
	"alpha/config"
	userDomain "alpha/domain/entity/admin/user"
	"go.uber.org/zap"
)

func ChangeStatus(uid uint64) error {
	userEntity := userDomain.NewEntity(uid)
	//更新信息
	if err := userEntity.ChangeStatus(); err != nil {
		config.Logger.Error("user change status",
			zap.Error(err),
		)
		return err
	}
	return nil
}
