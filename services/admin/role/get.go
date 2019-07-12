package role

import (
	"alpha/config"
	roleDomain "alpha/domain/entity/admin/role"
	"alpha/pkg/errno"
	"fmt"

	"go.uber.org/zap"
)

func Get(id uint64) (bool, error) {
	roleEntity := roleDomain.NewEntity(id)
	isNotFound, err := roleEntity.Info()
	if err != nil {
		config.Logger.Error("permission get",
			zap.Error(err),
		)
		return false, err
	}
	if isNotFound {
		return false, errno.ErrDBNotFoundRecord
	}
	fmt.Printf("%+v", roleEntity)
	return false, nil
}
