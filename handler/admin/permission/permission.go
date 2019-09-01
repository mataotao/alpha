package permission

type CreateRequest struct {
	Pid   uint64 `json:"pid"`
	Level uint8  `json:"level" valid:"required"`
	Data
}
type Data struct {
	Label     string `json:"label" valid:"required~请填写菜单名,stringlength(1|10)~菜单名最大长度10"`
	Sort      uint64 `json:"sort"`
	Url       string `json:"url" valid:"stringlength(1|500)~url最大长度500"`
	Cond      string `json:"cond" valid:"stringlength(1|2000)~条件最大长度2000"`
	Component string `json:"component" valid:"stringlength(1|2000)~组件最大长度2000"`
	Icon      string `json:"icon" valid:"stringlength(1|100)~icon最大长度100"`
}

type UpdateRequest struct {
	Id uint64 `uri:"id" valid:"required"`
	Data
}
type DeleteRequest struct {
	Id uint64 `uri:"id" valid:"required"`
}
type GetRequest struct {
	Id uint64 `uri:"id" valid:"required"`
}
