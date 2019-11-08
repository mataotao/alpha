package value_object

var RealtimeSkycon = map[string]string{
	"CLEAR_DAY":           "晴（白天）",
	"CLEAR_NIGHT":         "晴（夜间）",
	"PARTLY_CLOUDY_DAY":   "多云（白天）",
	"PARTLY_CLOUDY_NIGHT": "多云（夜间）",
	"CLOUDY":              "阴",
	"WIND":                "大风",
	"HAZE":                "雾霾",
	"RAIN":                "雨",
	"SNOW":                "雪",
}

type RealtimeStruct struct {
	Result RealtimeResult `json:"result"`
}
type RealtimeResult struct {
	Status      string  `json:"status"`      //	实况模块返回状态
	Temperature float64 `json:"temperature"` //	温度
	Skycon      string  `json:"skycon"`      //主要天气现象
	Comfort     Comfort `json:"comfort"`     //舒适度指数及其自然语言描述
	Pm25        float64 `json:"pm25"`        //pm25，质量浓度值
}

type Comfort struct {
	Index uint64 `json:"index"`
	Desc  string `json:"desc"`
}
