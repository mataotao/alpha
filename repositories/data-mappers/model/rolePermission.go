package model

type RolePermissionModel struct {
	BaseModel
	RoleId       int `json:"role_id"`
	PermissionId int `json:"permission_id"`
}

func (r *RolePermissionModel) TableName() string {
	return "role_permission"
}
