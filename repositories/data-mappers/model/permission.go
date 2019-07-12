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
type PermissionListInfo struct {
	Id            uint64                `json:"id"`
	Label         string                `json:"label" `
	IsContainMenu uint8                 `json:"is_contain_menu" `
	Pid           uint64                `json:"pid" `
	Level         uint8                 `json:"level" `
	Url           string                `json:"url" `
	Sort          uint64                `json:"sort" `
	Children      []*PermissionListInfo `json:"children"`
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

///获取一条
func (p *PermissionModel) Get(field string) (bool, error) {
	var isNotFound bool
	d := DB.Alpha.Select(field).First(&p)
	if d.RecordNotFound() {
		isNotFound = true
		return isNotFound, nil
	}
	if err := d.Error; err != nil {
		return isNotFound, err
	}

	return isNotFound, nil
}

func (p *PermissionModel) List(field string, ids []uint64) ([]*PermissionModel, error) {
	list := make([]*PermissionModel, 0)
	db := DB.Alpha.Select(field).Order("pid asc,sort desc,id asc")
	if len(ids) != 0 {
		db = db.Where("id in (?)", ids)
	}
	//查询数据
	if err := db.Find(&list).Error; err != nil {
		return list, err
	}
	return list, nil
}
func (p *PermissionModel) RecursivePermission(pid uint64, plist []*PermissionModel) []*PermissionListInfo {
	list := make([]*PermissionListInfo, 0)
	for i := range plist[:] {
		if pid == plist[i].Pid {
			cList := &PermissionListInfo{
				Id:            plist[i].Id,
				Label:         plist[i].Label,
				IsContainMenu: plist[i].IsContainMenu,
				Pid:           plist[i].Pid,
				Level:         plist[i].Level,
				Url:           plist[i].Url,
				Sort:          plist[i].Sort,
			}
			cList.Children = p.RecursivePermission(plist[i].Id, plist)
			list = append(list, cList)
		}
	}
	return list
}
