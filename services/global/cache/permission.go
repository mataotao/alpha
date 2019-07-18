package cache

import (
	"alpha/config"
	permissionDomain "alpha/domain/entity/admin/permission"
	"go.uber.org/zap"
)

func Permission(requestId string) error {
	permissionEntity := permissionDomain.NewEntityS(requestId)
	list, err := permissionEntity.All("*")
	if err != nil {
		config.Logger.Error("permission get",
			zap.Error(err),
		)
		return err
	}
	if err := permissionEntity.GenerateCache(list); err != nil {
		config.Logger.Error("permission get",
			zap.Error(err),
		)
		return err
	}
	return nil
}
