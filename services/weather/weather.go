package weather

import (
	"alpha/config"
	"alpha/domain/entity/weather"
	"go.uber.org/zap"
)

type weatherService struct {
}

func (w *weatherService) Send() {
	weatherEntity := weather.NewEntity(0)
	field := "name,phone,lng,lat,address,is_title"
	list, err := weatherEntity.SmsAll(field)
	if err != nil {
		config.Logger.Error("user list",
			zap.Error(err),
		)
		return
	}
	for i := range list[:] {
		weatherEntity.SetModel(list[i])
		if err := weatherEntity.Realtime(); err != nil {
			config.Logger.Error("user list",
				zap.Error(err),
			)
			return
		}
		if err := weatherEntity.Daily(); err != nil {
			config.Logger.Error("user list",
				zap.Error(err),
			)
			return
		}
		content := weatherEntity.Content()
		weatherEntity.Send(content)
	}
}
func NewWeatherService() *weatherService {
	s := new(weatherService)
	return s
}
