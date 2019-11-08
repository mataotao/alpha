package model

type SMSModel struct {
	BaseModel
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Lng     string `json:"lng"`
	Lat     string `json:"lat"`
	Address string `json:"address"`
	IsTitle int8   `json:"is_title"`
}

func (s *SMSModel) TableName() string {
	return "sms"
}

func (s *SMSModel) All(field string) ([]*SMSModel, error) {
	list := make([]*SMSModel, 0)
	db := DB.Alpha.Select(field)

	if err := db.Find(&list).Error; err != nil {
		return list, err
	}
	return list, nil
}
