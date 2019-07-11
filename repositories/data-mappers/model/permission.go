package model

type PermissionModel struct {
	BaseModel
	Label         string `json:"label"`
	IsContainMenu uint8  `json:"is_contain_menu"`
	Pid           uint64 `json:"pid"`
	Level         uint8  `json:"level"`
	Url           string `json:"url"`
	Sort          uint64 `json:"sort"`
	Cond          string `json:"cond"`
	Icon          string `json:"icon"`
}

//表名
func (p *PermissionModel) TableName() string {
	return "permission"
}

///创建
func (p *PermissionModel) Create() error {
	return DB.Alpha.Create(&p).Error
}

///修改
func (p *PermissionModel) Update() error {
	//new(PermissionModel) == &PermissionModel{}
	return DB.Alpha.Model(new(PermissionModel)).Updates(p).Error
}

///删除
func (p *PermissionModel) Delete() error {
	return DB.Alpha.Delete(&p).Error
}
