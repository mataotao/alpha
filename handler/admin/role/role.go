package role

type CreateRequest struct {
	Data
}
type DeleteRequest struct {
	Id uint64 `uri:"id" valid:"required"`
}
type GetRequest struct {
	Id uint64 `uri:"id" valid:"required"`
}
type UpdateRequest struct {
	Id uint64 `uri:"id" valid:"required"`
	Data
}
type Data struct {
	Name        string `json:"name" valid:"required~请填写角色名称"`
	Description string `json:"description" valid:"stringlength(1|500)~url最大长度500"`
	Permission  []int  `json:"permission" valid:"required~请选择权限"`
}
type ListRequest struct {
	Page  uint64 `form:"page"`
	Name  string `form:"name"`
	Limit uint64 `form:"limit"`
}
