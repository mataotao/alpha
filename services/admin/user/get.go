package user

import (
	"alpha/config"
	userDomain "alpha/domain/entity/admin/user"
	"alpha/pkg/errno"
	"time"

	"go.uber.org/zap"
)

type GetResponse struct {
	Id       uint64    `json:"id"`
	Username string    `json:"username"`
	Name     string    `json:"name"`
	Mobile   uint64    `json:"mobile"`
	Avatar   string    `json:"avatar"`
	LastTime time.Time `json:"last_time"`
	LastIp   string    `json:"last_ip"`
	IsRoot   byte      `json:"is_root"`
	Status   byte      `json:"status"`
	RoleIds  []uint64  `json:"role_ids"`
}

func Get(uid uint64) (*GetResponse, error) {
	getResponse := new(GetResponse)
	userEntity := userDomain.NewEntity(uid)
	notfound, err := userEntity.Get("*")
	if err != nil {
		config.Logger.Error("user get",
			zap.Error(err),
		)
		return getResponse, err
	}
	if notfound {
		return getResponse, errno.ErrDBNotFoundRecord
	}

	if _, err := userEntity.GetRoleIds(); err != nil {
		config.Logger.Error("user user_permission_ids",
			zap.Error(err),
		)
		return getResponse, err
	}
	getResponse = &GetResponse{
		Id:       userEntity.UserModel.Id,
		Username: userEntity.UserModel.Username,
		Name:     userEntity.UserModel.Name,
		Mobile:   userEntity.UserModel.Mobile,
		Avatar:   userEntity.UserModel.Avatar,
		LastTime: userEntity.UserModel.LastTime,
		LastIp:   userEntity.UserModel.LastIp,
		IsRoot:   userEntity.UserModel.IsRoot,
		Status:   userEntity.UserModel.Status,
		RoleIds:  userEntity.RoleIds,
	}

	return getResponse, nil
}
