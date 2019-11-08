package queue

import (
	services "alpha/services/weather"
	"fmt"
	"time"
)

func weather() {
	timer := time.NewTimer(0)
	fmt.Printf("queue:weather:%v SUCCESS \n", NextDayTime()+6*time.Hour)
	for {
		if !timer.Stop() {
			select {
			case <-timer.C:
			default:
			}
		}
		timer.Reset(NextDayTime() + 6*time.Hour)
		select {
		case <-timer.C:
			weatherService := services.NewWeatherService()
			weatherService.Send()
			continue
		}
	}

}
func NextDayTime() time.Duration {
	now := time.Now()
	next := now.Add(time.Hour * 24)
	next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
	return next.Sub(now)
}
