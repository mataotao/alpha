package permission

import (
	"alpha/domain/entity"
	"alpha/pkg/constvar"
	redis "alpha/repositories/data-mappers/go-redis"
	"alpha/repositories/data-mappers/model"

	"fmt"
	"regexp"
	"strings"
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
	//删除旧权限
	keys, err := redis.Scan(fmt.Sprintf(CachePermissionKey, "*"), constvar.DefaultScanCount)
	if err != nil {
		return err
	}
	if len(keys) > 0 {
		pipe.Del(keys...)
	}
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
		for j := range conditions[:] {
			k := fmt.Sprintf(CachePermissionKey, conditions[j])
			pipe.Set(k, list[i].Id, 0)
		}
	}
	_, err = pipe.Exec()
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
