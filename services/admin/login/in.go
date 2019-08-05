package login

import (
	"alpha/config"
	userDomain "alpha/domain/entity/admin/user"
	"alpha/pkg/errno"
	userService "alpha/services/admin/user"
	"go.uber.org/zap"
)

type InResponse struct {
	Token  string `json:"token"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

func In(username, pwd, ip string) (*InResponse, error) {
	inResponse := new(InResponse)
	userEntity := userDomain.NewEntityS(username)
	notfound, err := userEntity.Get("*")
	if err != nil {
		config.Logger.Error("login in",
			zap.Error(err),
		)
		return inResponse, err
	}
	if notfound {
		return inResponse, errno.ErrDBNotFoundRecord
	}
	//检测密码
	if err := userEntity.Compare(pwd); err != nil {
		config.Logger.Error("login in",
			zap.Error(err),
		)
		return inResponse, errno.ErrUserNameOrPassword
	}
	//检测用户冻结
	if userEntity.IsFreeze() {
		return inResponse, errno.ErrUserFreeze
	}
	//获取权限id
	if _, err := userService.PermissionIds(userEntity); err != nil {
		config.Logger.Error("login in",
			zap.Error(err),
		)
		return inResponse, err
	}
	//权限id写入redis 位图
	if err := userEntity.SetPermissionToCache(); err != nil {
		config.Logger.Error("login in",
			zap.Error(err),
		)
		return inResponse, err
	}
	userEntity.UserModel.LastIp = ip
	//更新登录信息
	if err := userEntity.UpdateLogin(); err != nil {
		config.Logger.Error("login in",
			zap.Error(err),
		)
		return inResponse, err
	}
	//生成token
	token, err := userEntity.TokenSign()
	if err != nil {
		config.Logger.Error("login in",
			zap.Error(err),
		)
		return inResponse, err
	}
	inResponse.Name = userEntity.UserModel.Name
	inResponse.Avatar = userEntity.UserModel.Avatar
	inResponse.Token = token
	return inResponse, nil
}
