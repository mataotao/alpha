package model

import (
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
	var isNotFound bool
	d := DB.Alpha.Select(field).First(&r)
	if d.RecordNotFound() {
		isNotFound = true
		return isNotFound, nil
	}
	if err := d.Error; err != nil {
		return isNotFound, err
	}

	return isNotFound, nil
}
