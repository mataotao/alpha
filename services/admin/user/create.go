package user

import (
	"alpha/config"
	userDomain "alpha/domain/entity/admin/user"
	"alpha/pkg/errno"
	"alpha/repositories/data-mappers/model"

	"go.uber.org/zap"
)

func Create(user *model.UserModel, roleIds []uint64) error {
	userEntity := userDomain.NewEntity(0)
	userEntity.UserModel = *user
	//检查用户名唯一
	unique, err := userEntity.Unique()
	if err != nil {
		config.Logger.Error("user create",
			zap.Error(err),
		)
		return err
	}
	if unique == false {
		return errno.ErrUserNameNotUnique
	}
	//加密密码
	if err := userEntity.Encrypt(); err != nil {
		config.Logger.Error("user create",
			zap.Error(err),
		)
		return err

	}
	if err := userEntity.Create(roleIds); err != nil {
		config.Logger.Error("user create",
			zap.Error(err),
		)
		return err
	}
	return nil
}

//用户拥有的权限
func PermissionIds(userEntity *userDomain.Entity) ([]uint64, error) {
	ids := make([]uint64, 0)
	if userEntity.IsRoot() {
		permissionModel := new(model.PermissionModel)
		plist, err := permissionModel.All("id")
		if err != nil {
			config.Logger.Error("user user_permission_ids",
				zap.Error(err),
			)
			return ids, err
		}
		ids = permissionModel.Ids(plist)
	} else {
		_, err := userEntity.GetRoleIds()
		if err != nil {
			config.Logger.Error("user user_permission_ids",
				zap.Error(err),
			)
			return ids, err
		}
		ids, err = userEntity.GetPermissionIds(userEntity.RoleIds)
		if err != nil {
			config.Logger.Error("user user_permission_ids",
				zap.Error(err),
			)
			return ids, err
		}
	}
	return ids, nil
}
