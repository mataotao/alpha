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
