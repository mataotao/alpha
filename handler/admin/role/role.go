package role

type CreateRequest struct {
	Name        string `json:"name" valid:"required"`
	Description string `json:"description" valid:"stringlength(1|500)~url最大长度500"`
	Permission  []int  `json:"permission" valid:"required"`
}
