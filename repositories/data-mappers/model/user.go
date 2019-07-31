package model

import (
	"alpha/pkg/constvar"

	"github.com/jinzhu/gorm"

	"fmt"
	"strings"
	"time"
)

const (
	ON     byte = iota + 1
	FREEZE      //冻结
)

type UserModel struct {
	BaseModel
	Username string    `json:"username"`
	Name     string    `json:"name"`
	Mobile   uint64    `json:"mobile"`
	Password string    `json:"password"`
	HeadImg  string    `json:"head_img"`
	LastTime time.Time `json:"last_time"`
	LastIp   string    `json:"last_ip"`
	IsRoot   uint8     `json:"is_root"`
	Status   uint8     `json:"status"`
}

func (u *UserModel) TableName() string {
	return "user"
}

func (u *UserModel) Get(field string) (bool, error) {
	var notFound bool
	db := DB.Alpha.Select(field)
	if u.Username != "" {
		db = db.Where("username = ?", u.Username)
	}
	d := db.First(&u)

	if d.RecordNotFound() {
		notFound = true
		return notFound, nil
	}
	if err := d.Error; err != nil {
		return notFound, err
	}
	return notFound, nil
}

func (u *UserModel) Create(roleIds []uint64) error {
	tx := DB.Alpha.Begin()
	if err := tx.Create(&u).Error; err != nil {
		tx.Rollback()
		return err
	}

	var valueSql strings.Builder
	for i := range roleIds[:] {
		valueSql.WriteString(fmt.Sprintf("(%d,%d),", u.Id, roleIds[i]))
	}
	//转化成字符串
	sql := fmt.Sprintf(`INSERT INTO %s (user_id,role_id) VALUES %s`, new(UserRoleModel).TableName(), strings.TrimSuffix(valueSql.String(), ","))
	if err := tx.Exec(sql).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil

}

func (u *UserModel) Update(roleIds []uint64) error {
	tx := DB.Alpha.Begin()
	if err := tx.Model(u).Update(u).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where("user_id = ?", u.Id).Delete(UserRoleModel{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	var valueSql strings.Builder
	for i := range roleIds[:] {
		valueSql.WriteString(fmt.Sprintf("(%d,%d),", u.Id, roleIds[i]))
	}
	//转化成字符串
	sql := fmt.Sprintf(`INSERT INTO %s (user_id,role_id) VALUES %s`, new(UserRoleModel).TableName(), strings.TrimSuffix(valueSql.String(), ","))
	if err := tx.Exec(sql).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil

}
func (u *UserModel) ChangeStatus() error {
	cond := "CASE `status` WHEN ? THEN ? WHEN ? THEN ? END"
	return DB.Alpha.Model(u).Update("status", gorm.Expr(cond, FREEZE, ON, ON, FREEZE)).Error
}

func (u *UserModel) Updates(data map[string]interface{}) error {
	return DB.Alpha.Model(&UserModel{}).Where("id = ?", u.Id).Updates(data).Error
}
func (u *UserModel) List(field string, page, limit uint64) ([]*UserModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}
	if page == 0 {
		page = constvar.DefaultPage
	}
	start := (page - 1) * limit
	var count uint64
	list := make([]*UserModel, 0)
	db := DB.Alpha.Select(field)
	if u.Name != "" {
		db = db.Where(fmt.Sprintf("name like '%%%s%%'", u.Name))
	}
	if u.Username != "" {
		db = db.Where(fmt.Sprintf("username like '%%%s%%'", u.Username))
	}
	if u.Mobile != 0 {
		db = db.Where(fmt.Sprintf("mobile like '%%%d%%'", u.Mobile))
	}

	if err := db.Model(&UserModel{}).Count(&count).Error; err != nil {
		return list, count, err
	}

	if err := db.Offset(start).Limit(limit).Order("id desc").Find(&list).Error; err != nil {
		return list, count, err
	}
	return list, count, nil
}
