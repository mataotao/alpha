package user

type CreateRequest struct {
	Username string   `json:"username" valid:"required,stringlength(1|20)~密码最大长度20"`
	Name     string   `json:"name" valid:"required,stringlength(1|20)~密码最大长度20"`
	Mobile   uint64   `json:"mobile" valid:"required,length(11|11)"`
	Password string   `json:"password" valid:"required,stringlength(1|20)~密码最大长度20"`
	HeadImg  string   `json:"head_img" valid:"stringlength(1|500)~密码最大长度500"`
	RoleIds  []uint64 `json:"role_ids" valid:"required"`
}
