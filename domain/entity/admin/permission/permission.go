package permission

import (
	"alpha/domain/entity"
	"alpha/repositories/data-mappers/model"
)

type Entity struct {
	entity.Entity
	model.PermissionModel
}

func NewEntity(id uint64) *Entity {
	e := new(Entity)
	e.Entity.SetId(id)
	return e
}
