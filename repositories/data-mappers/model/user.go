package model

import (
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
	if u.Id != 0 {
		db = db.Where("id = ?", u.Id)
	}
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
func (u *UserModel) Updates(data map[string]interface{}) error {
	return DB.Alpha.Model(&UserModel{}).Where("id = ?", u.Id).Updates(data).Error
}
