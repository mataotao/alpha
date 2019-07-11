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
