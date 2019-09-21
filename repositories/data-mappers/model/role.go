package model

import (
	"alpha/pkg/constvar"

	"fmt"
	"strings"
)

type RoleModel struct {
	BaseModel
	Name        string `json:"name"`
	Description string `json:"description"`
}
type RoleInfo struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Permission  []int  `json:"permission"`
}

func (r *RoleModel) TableName() string {
	return "role"
}

func (r *RoleModel) Create(p []int) error {
	//开启事务
	tx := DB.Alpha.Begin()
	if err := tx.Create(&r).Error; err != nil {
		tx.Rollback()
		return err
	}

	var valueSql strings.Builder
	for i := range p[:] {
		valueSql.WriteString(fmt.Sprintf("(%d,%d),", r.Id, p[i]))
	}
	//转化成字符串
	sql := fmt.Sprintf(`INSERT INTO %s (role_id,permission_id) VALUES %s`, new(RolePermissionModel).TableName(), strings.TrimSuffix(valueSql.String(), ","))
	if err := tx.Exec(sql).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (r *RoleModel) Delete() error {
	//开启事务
	tx := DB.Alpha.Begin()
	if err := tx.Delete(&r).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Where("role_id = ?", r.Id).Delete(RolePermissionModel{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
func (r *RoleModel) Update(p []int) error {
	tx := DB.Alpha.Begin()
	if err := tx.Model(&r).Updates(r).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Where("role_id = ?", r.Id).Delete(RolePermissionModel{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	var valueSql strings.Builder
	for i := range p[:] {
		valueSql.WriteString(fmt.Sprintf("(%d,%d),", r.Id, p[i]))
	}
	sql := fmt.Sprintf(`INSERT INTO %s (role_id,permission_id) VALUES %s`, new(RolePermissionModel).TableName(), strings.TrimSuffix(valueSql.String(), ","))
	if err := tx.Exec(sql).Error; err != nil {
		return err
	}
	tx.Commit()
	return nil

}

func (r *RoleModel) Get(field string) (bool, error) {
	var notFound bool
	d := DB.Alpha.Select(field).First(&r)
	if d.RecordNotFound() {
		notFound = true
		return notFound, nil
	}
	if err := d.Error; err != nil {
		return notFound, err
	}

	return notFound, nil
}
func (r *RoleModel) List(field string, page uint64, limit uint64) ([]*RoleModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}
	if page == 0 {
		page = constvar.DefaultPage
	}
	start := (page - 1) * limit
	var count uint64
	list := make([]*RoleModel, 0)
	db := DB.Alpha.Select(field)
	if r.Name != "" {
		db = db.Where(fmt.Sprintf("name like '%%%s%%'", r.Name))
	}

	if err := db.Model(&RoleModel{}).Count(&count).Error; err != nil {
		return list, count, err
	}

	if err := db.Offset(start).Limit(limit).Order("id desc").Find(&list).Error; err != nil {
		return list, count, err
	}
	return list, count, nil

}
func (r *RoleModel) All(field string) ([]*RoleModel, error) {
	list := make([]*RoleModel, 0)
	db := DB.Alpha.Select(field)
	if err := db.Order("id desc").Find(&list).Error; err != nil {
		return list, err
	}
	return list, nil

}
