package queue

import (
	services "alpha/services/weather"
)

func weather() {
	weatherService := services.NewWeatherService()
	weatherService.Send()
}
