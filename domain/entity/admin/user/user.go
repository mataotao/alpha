package user

import (
	"alpha/domain/entity"
	"alpha/pkg/auth"
	"alpha/pkg/token"
	redis "alpha/repositories/data-mappers/go-redis"
	"alpha/repositories/data-mappers/model"
	sliceUtil "alpha/repositories/util/slice"

	"fmt"
	"time"
)

const (
	PermissionKey = "alpha:user:permission:%d"
)
const (
	PermissionOff = iota
	PermissionOn
)

type Entity struct {
	entity.Entity
	model.UserModel
	RoleIds       []uint64
	PermissionIds []uint64
}

//创建用户
func (e *Entity) Create(roleIds []uint64) error {
	if err := (&e.UserModel).Create(roleIds); err != nil {
		return err
	}
	(&e.Entity).SetId(e.UserModel.Id)
	return nil
}

//更新用户
func (e *Entity) Update(roleIds []uint64) error {
	return (&e.UserModel).Update(roleIds)
}

//检查用户名唯一
func (e *Entity) Unique() (bool, error) {
	e.BindSId()
	notFound, err := (&e.UserModel).Get("id")
	if err != nil {
		return false, err
	}
	if !notFound {
		return false, nil
	}
	return true, nil
}

//加密
func (e *Entity) Encrypt() (err error) {
	e.UserModel.Password, err = auth.Encrypt(e.UserModel.Password)
	return
}

//获取信息
func (e *Entity) Get(field string) (bool, error) {
	e.BindIds()
	notFound, err := (&e.UserModel).Get(field)
	if err != nil {
		return notFound, err
	}

	return notFound, nil
}

//是否冻结
func (e *Entity) IsFreeze() bool {
	return e.UserModel.Status == model.FREEZE
}

//检测密码
func (e *Entity) Compare(pwd string) error {
	return auth.Compare(e.UserModel.Password, pwd)
}

//是否是超级管理员
func (e *Entity) IsRoot() bool {
	return e.UserModel.IsRoot == model.ON
}

//获取用户ids
func (e *Entity) GetRoleIds() ([]uint64, error) {
	var ids []uint64
	e.BindId()
	userRoleModel := &model.UserRoleModel{
		UserId: e.UserModel.Id,
	}
	list, err := userRoleModel.AllByUserId("role_id")
	if err != nil {
		return ids, err
	}
	ids = userRoleModel.RoleIds(list)
	e.RoleIds = ids
	return ids, nil
}

//获取权限ids
func (e *Entity) GetPermissionIds(roleIds []uint64) ([]uint64, error) {
	var ids []uint64
	rolePermissionModel := new(model.RolePermissionModel)
	list, err := rolePermissionModel.AllByRoleIds("permission_id", roleIds)
	if err != nil {
		return ids, err
	}
	ids = rolePermissionModel.PermissionIds(list)
	ids = sliceUtil.RemoveDuplicateElementUint64(ids)
	e.PermissionIds = ids
	return ids, nil
}

//存入权限到缓存
func (e *Entity) SetPermissionToCache() error {
	e.BindId()
	if len(e.PermissionIds) == 0 {
		panic("permission_ids为空")
	}
	k := fmt.Sprintf(PermissionKey, e.UserModel.Id)
	pipe := redis.Client.Client.TxPipeline()
	pipe.Del(k)
	for i := range e.PermissionIds[:] {
		pipe.SetBit(k, int64(e.PermissionIds[i]), PermissionOn)
	}
	_, err := pipe.Exec()
	if err != nil {
		return err
	}
	return nil
}

//更新登录信息
func (e *Entity) UpdateLogin() error {
	data := make(map[string]interface{})
	data["last_ip"] = e.UserModel.LastIp
	data["last_time"] = time.Now()
	err := (&e.UserModel).Updates(data)
	if err != nil {
		return err
	}
	return nil
}

//更新密码
func (e *Entity) UpdatePwd() error {
	data := make(map[string]interface{})
	data["password"] = e.UserModel.Password
	err := (&e.UserModel).Updates(data)
	if err != nil {
		return err
	}
	return nil
}

//改变状态
func (e *Entity) ChangeStatus() error {
	e.BindId()
	err := (&e.UserModel).ChangeStatus()
	if err != nil {
		return err
	}
	return nil
}

//生成登录token
func (e *Entity) TokenSign() (string, error) {
	t, err := token.Sign(token.Context{ID: e.UserModel.Id, Username: e.UserModel.Username}, "")
	if err != nil {
		return t, err
	}
	return t, nil
}

//领域检查
func (e *Entity) BindIds() {
	if id := (&e.Entity).GetId(); id != 0 {
		e.UserModel.Id = id
	}
	if sid := (&e.Entity).GetSId(); sid != "" {
		e.UserModel.Username = sid
	}
	if e.UserModel.Id == 0 && e.UserModel.Username == "" {
		panic("请传入唯一标识")
	}
}
func (e *Entity) BindId() {
	if id := (&e.Entity).GetId(); id != 0 {
		e.UserModel.Id = id
	}
	if e.UserModel.Id == 0 {
		panic("请传入唯一标识")
	}
}
func (e *Entity) BindSId() {
	if sid := (&e.Entity).GetSId(); sid != "" {
		e.UserModel.Username = sid
	}
	if e.UserModel.Username == "" {
		panic("请传入唯一标识")
	}
}
func NewEntity(id uint64) *Entity {
	e := new(Entity)
	(&e.Entity).SetId(id)
	return e
}

func NewEntityS(sid string) *Entity {
	e := new(Entity)
	(&e.Entity).SetSId(sid)
	return e
}
