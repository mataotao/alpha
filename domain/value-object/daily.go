package value_object

type DailyStruct struct {
	Result struct {
		Daily struct {
			Temperature []dailyTemperature `json:"temperature"` //温度，最大值，平均值，最小值
			Skycon      []dailySkycon      `json:"skycon"`      //全天，主要天气现象
			ColdRisk    []dailyColdRisk    `json:"coldRisk"`    //全天，感冒指数和自然语言描述
			Comfort     []dailyColdRisk    `json:"comfort"`     //全天，舒适度指数和自然语言描述
			Pm25        []dailyTemperature `json:"pm25"`        //PM2.5，最大值，平均值最小值
		} `json:"daily"`
	} `json:"result"`
}
type dailyTemperature struct {
	Date string  `json:"date"`
	Max  float64 `json:"max"`
	Avg  float64 `json:"avg"`
	Min  float64 `json:"min"`
}
type dailySkycon struct {
	Date  string `json:"date"`
	Value string `json:"value"`
}
type dailyColdRisk struct {
	Desc     string `json:"desc"`
	Datetime string `json:"datetime"`
}
