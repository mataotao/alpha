package permission

import (
	"alpha/config"
	"alpha/repositories/data-mappers/model"
	"go.uber.org/zap"
)

type ListInfo struct {
	Id            uint64      `json:"id"`
	Label         string      `json:"label" `
	IsContainMenu uint8       `json:"is_contain_menu" `
	Pid           uint64      `json:"pid" `
	Level         uint8       `json:"level" `
	Url           string      `json:"url" `
	Sort          uint64      `json:"sort" `
	Children      []*ListInfo `json:"children"`
}

func List() ([]*ListInfo, error) {
	list := make([]*ListInfo, 0)
	//TODO 用户权限过滤
	//获取权限列表
	p := new(model.PermissionModel)
	plist, err := p.List("*", []uint64{})
	if err != nil {
		config.Logger.Error("permission get",
			zap.Error(err),
		)
		return list, err
	}
	list = RecursivePermission(0, plist)
	return list, nil
}
func RecursivePermission(pid uint64, plist []*model.PermissionModel) []*ListInfo {
	list := make([]*ListInfo, 0)
	for i := range plist[:] {
		if pid == plist[i].Pid {
			cList := &ListInfo{
				Id:            plist[i].Id,
				Label:         plist[i].Label,
				IsContainMenu: plist[i].IsContainMenu,
				Pid:           plist[i].Pid,
				Level:         plist[i].Level,
				Url:           plist[i].Url,
				Sort:          plist[i].Sort,
			}
			cList.Children = RecursivePermission(plist[i].Id, plist)
			list = append(list, cList)
		}
	}
	return list
}
