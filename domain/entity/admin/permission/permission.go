package permission

import (
	"alpha/domain/entity"
	"fmt"
	"regexp"
	"strings"

	redis "alpha/repositories/data-mappers/go-redis"
	"alpha/repositories/data-mappers/model"
)

const (
	CachePermissionKey = "alpha:cache:permission:%s"
)

type Entity struct {
	entity.Entity
	model.PermissionModel
}

func (e *Entity) All(field string) ([]*model.PermissionModel, error) {
	list, err := (&e.PermissionModel).All(field)
	if err != nil {
		return list, err
	}
	return list, nil
}
func (e *Entity) GenerateCache(list []*model.PermissionModel) error {
	pipe := redis.Client.Client.TxPipeline()
	for i := range list[:] {
		//判断是否为空
		matched, err := regexp.Match("^[\\s\\S]*.*[^\\s][\\s\\S]*$", []byte(list[i].Cond))
		if err != nil {
			return err
		}
		//为空就跳过
		if !matched {
			continue
		}
		conditions := strings.Split(list[i].Cond, ",")
		for i := range conditions[:] {
			k := fmt.Sprintf(CachePermissionKey, conditions[i])
			pipe.Set(k, list[i].Id, 0)
		}
	}
	_, err := pipe.Exec()
	if err != nil {
		return err
	}
	return nil
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
