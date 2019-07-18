package role

import (
	"alpha/domain/entity"
	"alpha/repositories/data-mappers/model"
)

type Entity struct {
	entity.Entity
	model.RoleModel
}

func (e *Entity) Info(field string) (bool, error) {
	var m model.RoleModel
	m.Id = e.Entity.GetId()
	isNotFound, err := m.Get(field)
	if err != nil {
		return isNotFound, err
	}
	e.RoleModel = m
	return isNotFound, nil
}
func NewEntity(id uint64) *Entity {
	e := new(Entity)
	(&e.Entity).SetId(id)
	return e
}
