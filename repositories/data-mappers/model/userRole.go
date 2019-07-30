package model

type UserRoleModel struct {
	BaseModel
	UserId uint64 `json:"user_id"`
	RoleId uint64 `json:"role_id"`
}

func (u *UserRoleModel) TableName() string {
	return "user_role"
}
func (u *UserRoleModel) AllByUserId(field string) ([]*UserRoleModel, error) {
	list := make([]*UserRoleModel, 0)
	db := DB.Alpha.Select(field)
	if u.UserId != 0 {
		db = db.Where("user_id = ?", u.UserId)
	}
	//查询数据
	if err := db.Find(&list).Error; err != nil {
		return list, err
	}
	return list, nil
}
func (u *UserRoleModel) RoleIds(list []*UserRoleModel) []uint64 {
	ids := make([]uint64, 0, len(list))
	for i := range list[:] {
		ids = append(ids, list[i].RoleId)
	}
	return ids
}
