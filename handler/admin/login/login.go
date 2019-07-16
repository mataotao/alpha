package login

type InRequest struct {
	Username string `json:"username" valid:"required~请填写用户名"`
	Password string `json:"password" valid:"required~请填写密码"`
}
