package user

import (
	"alpha/domain/entity"
	"alpha/repositories/data-mappers/model"
)

type Entity struct {
	entity.Entity
	model.UserModel
}

//创建用户
func (e *Entity) Create(roleIds []uint64) error {
	if err := e.UserModel.Create(roleIds); err != nil {
		return err
	}
	e.Entity.SetId(e.UserModel.Id)
	return nil
}

//检查用户名唯一
func (e *Entity) Unique() bool {
	if e.UserModel.Username == "" {
		return false
	}
	return e.UserModel.Unique()
}

//加密
func (e *Entity) Encrypt() error {
	return e.UserModel.Encrypt()
}
func NewEntity(id uint64) *Entity {
	e := new(Entity)
	e.Entity.SetId(id)
	return e
}
