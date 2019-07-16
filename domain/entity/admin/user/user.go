package user

import (
	"alpha/domain/entity"
	"alpha/pkg/auth"
	redis "alpha/repositories/data-mappers/go-redis"
	"alpha/repositories/data-mappers/model"
	sliceUtil "alpha/repositories/util/slice"
	"fmt"
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
	e.Entity.SetId(e.UserModel.Id)
	return nil
}

//检查用户名唯一
func (e *Entity) Unique() (bool, error) {
	if e.UserModel.Username == "" {
		panic("username为空")
	}
	notFound, err := (&e.UserModel).Get("id")
	if err != nil {
		return false, err
	}
	if notFound == false {
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
func (e *Entity) Get() (bool, error) {
	if id := e.Entity.GetId(); id != 0 {
		e.UserModel.Id = id
	}
	if sid := e.Entity.GetSId(); sid != "" {
		e.UserModel.Username = sid
	}
	notFound, err := (&e.UserModel).Get("*")
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

//获取用户权限ids
func (e *Entity) GetRoleIds() ([]uint64, error) {
	var ids []uint64
	if e.UserModel.Id == 0 {
		panic("id为空")
	}
	userRoleModel := &model.UserRoleModel{
		UserId: e.UserModel.Id,
	}
	list, err := userRoleModel.AllByUserId("id")
	if err != nil {
		return ids, err
	}
	ids = userRoleModel.Ids(list)
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
func (e *Entity) SetPermissionToCache() error {
	if e.UserModel.Id == 0 {
		panic("id为空")
	}
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
func NewEntity(id uint64) *Entity {
	e := new(Entity)
	e.Entity.SetId(id)
	return e
}

func NewEntityS(sid string) *Entity {
	e := new(Entity)
	e.Entity.SetSId(sid)
	return e
}
