package role

import (
	"alpha/config"
	"alpha/repositories/data-mappers/model"
	"go.uber.org/zap"
)

func All() ([]*model.RoleModel, error) {
	roleModel := new(model.RoleModel)
	list, err := roleModel.All("*")
	if err != nil {
		config.Logger.Error("role all",
			zap.Error(err),
		)
		return list, err
	}
	return list, nil
}
