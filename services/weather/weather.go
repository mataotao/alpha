package weather

import (
	"alpha/domain/entity/weather"
)

type weatherService struct {
}

func (w *weatherService) Send() {
	weatherEntity := weather.NewEntity(0)
	content := weatherEntity.Content()
	weatherEntity.Send("+8617628293814", content)
}
func NewWeatherService() *weatherService {
	s := new(weatherService)
	return s
}
