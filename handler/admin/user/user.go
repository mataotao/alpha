package user

type CreateRequest struct {
	Username string `json:"username" valid:"required,stringlength(1|20)~用户名最大长度20"`
	PasswordInfo
	Info
}
type UpdateRequest struct {
	Id uint64 `uri:"id" valid:"required"`
	Info
}
type PasswordInfo struct {
	Password string `json:"password" valid:"required,stringlength(1|20)~密码最大长度20"`
}
type PwdRequest struct {
	Id uint64 `uri:"id" valid:"required"`
	PasswordInfo
}
type StatusRequest struct {
	Id uint64 `uri:"id" valid:"required"`
}
type GetRequest struct {
	Id uint64 `uri:"id" valid:"required"`
}
type ListRequest struct {
	Page     uint64 `form:"page"`
	Limit    uint64 `form:"limit"`
	Username string `form:"username"`
	Name     string `form:"name"`
	Mobile   uint64 `form:"mobile"`
}

type Info struct {
	Name    string   `json:"name" valid:"required,stringlength(1|20)~名称最大长度20"`
	Mobile  uint64   `json:"mobile" valid:"required~请填写手机号,length(11|11)~手机号格式不正确"`
	Avatar  string   `json:"avatar" valid:"stringlength(1|500)~头像最大长度500"`
	RoleIds []uint64 `json:"role_ids" valid:"required~请选择角色"`
}
