package model

type RolePermissionModel struct {
	BaseModel
	RoleId       uint64 `json:"role_id"`
	PermissionId uint64 `json:"permission_id"`
}

func (rp *RolePermissionModel) TableName() string {
	return "role_permission"
}
func (rp *RolePermissionModel) All(field string) ([]*RolePermissionModel, error) {
	list := make([]*RolePermissionModel, 0)
	db := DB.Alpha.Select(field)
	if rp.RoleId != 0 {
		db = db.Where("role_id = ?", rp.RoleId)
	}
	//查询数据
	if err := db.Find(&list).Error; err != nil {
		return list, err
	}
	return list, nil
}
func (rp *RolePermissionModel) PermissionIds(list []*RolePermissionModel) []uint64 {
	permissionIds := make([]uint64, 0, len(list))
	for i := range list[:] {
		permissionIds = append(permissionIds, list[i].PermissionId)
	}
	return permissionIds
}
