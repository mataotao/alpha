package model

type UserRoleModel struct {
	BaseModel
	UserId uint64 `json:"user_id"`
	RoleId uint64 `json:"role_id"`
}

func (u *UserRoleModel) TableName() string {
	return "user_role"
}
