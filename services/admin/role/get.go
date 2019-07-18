package role

import (
	"alpha/config"
	roleDomain "alpha/domain/entity/admin/role"
	"alpha/pkg/errno"
	"alpha/repositories/data-mappers/model"
	"go.uber.org/zap"
)

type Info struct {
	Id           uint64   `json:"id"`
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	PermissionId []uint64 `json:"permission_id"`
}

func Get(id uint64) (*Info, error) {
	info := new(Info)
	roleEntity := roleDomain.NewEntity(id)
	notFound, err := roleEntity.Info("*")
	if err != nil {
		config.Logger.Error("role get",
			zap.Error(err),
		)
		return info, err
	}
	if notFound {
		return info, errno.ErrDBNotFoundRecord
	}
	rp := new(model.RolePermissionModel)
	rp.RoleId = id
	permissionList, err := rp.All("*")
	if err != nil {
		config.Logger.Error("role get",
			zap.Error(err),
		)
		return info, err
	}

	info.Id = roleEntity.Id
	info.Name = roleEntity.Name
	info.Description = roleEntity.Description
	info.PermissionId = rp.PermissionIds(permissionList)

	return info, nil
}
