package role

import (
	"alpha/config"
	"alpha/repositories/data-mappers/model"

	"go.uber.org/zap"
)

type ListResponse struct {
	List  []*model.RoleModel `json:"list"`
	Count uint64             `json:"count"`
}

func List(name string, page, limit uint64) (*ListResponse, error) {
	list := new(ListResponse)
	var err error
	roleModel := &model.RoleModel{
		Name: name,
	}
	list.List, list.Count, err = roleModel.List("*", page, limit)
	if err != nil {
		config.Logger.Error("role list",
			zap.Error(err),
		)
		return list, err
	}
	return list, nil
}
