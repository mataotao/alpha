package permission

import (
	"alpha/config"
	"alpha/pkg/errno"
	"alpha/repositories/data-mappers/model"

	"go.uber.org/zap"
)

func Get(id uint64) (*model.PermissionModel, error) {
	p := new(model.PermissionModel)
	p.Id = id
	isNotFound, err := p.Get("*")
	if err != nil {
		config.Logger.Error("permission get",
			zap.Error(err),
		)
		return p, err
	}
	if isNotFound {
		return p, errno.ErrDBNotFoundRecord
	}
	return p, nil
}
