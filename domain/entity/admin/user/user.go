package user

import (
	"alpha/domain/entity"
	"alpha/pkg/auth"
	"alpha/repositories/data-mappers/model"
)

type Entity struct {
	entity.Entity
	model.UserModel
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
		return false, nil
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
