package permission

type CreateRequest struct {
	Pid   uint64 `json:"pid"`
	Level uint8  `json:"level" valid:"required"`
	DataRequest
}
type DataRequest struct {
	Label         string `json:"label" valid:"required,stringlength(1|10)~菜单名最大长度10"`
	Sort          uint64 `json:"sort"`
	IsContainMenu uint8  `json:"is_contain_menu" valid:"in(1,2)~此类型不存在"`
	Url           string `json:"url" valid:"required,stringlength(1|500)~url最大长度500"`
	Cond          string `json:"cond" valid:"stringlength(1|2000)~条件最大长度2000"`
	Icon          string `json:"icon" valid:"stringlength(1|100)~icon最大长度100"`
}

type UpdateRequest struct {
	Id uint64 `uri:"id" valid:"required"`
	DataRequest
}
type DeleteRequest struct {
	Id uint64 `uri:"id" valid:"required"`
}
