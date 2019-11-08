package weather

import (
	"alpha/config"
	baseEntity "alpha/domain/entity"
	valueObject "alpha/domain/value-object"
	"alpha/repositories/data-mappers/model"

	resty "github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const (
	_ int8 = iota
	ON
	OFF
)

type entity struct {
	baseEntity.Entity
	valueObject.RealtimeStruct
	valueObject.DailyStruct
	model.SMSModel
}

func (e *entity) Realtime() error {
	token := viper.GetString("colorful.token")
	realUrl := fmt.Sprintf("https://api.caiyunapp.com/v2/%s/%s,%s/realtime.json", token, e.SMSModel.Lng, e.SMSModel.Lat)
	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		Get(realUrl)
	config.Logger.Info("Realtime",
		zap.Any("url", realUrl),
		zap.Any("res", resp),
		zap.Error(err),
	)
	if err != nil {
		return err
	}
	if err := json.Unmarshal([]byte(resp.String()), &e.RealtimeStruct); err != nil {
		return err
	}

	return nil
}
func (e *entity) Daily() error {
	token := viper.GetString("colorful.token")
	realUrl := fmt.Sprintf("https://api.caiyunapp.com/v2/%s/%s,%s/daily.json", token, e.SMSModel.Lng, e.SMSModel.Lat)
	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		Get(realUrl)
	config.Logger.Info("Daily",
		zap.Any("url", realUrl),
		zap.Any("res", resp),
		zap.Error(err),
	)
	if err != nil {
		return err
	}
	if err := json.Unmarshal([]byte(resp.String()), &e.DailyStruct); err != nil {
		return err
	}

	return nil
}

func (e *entity) Content() string {
	var s strings.Builder
	if e.SMSModel.IsTitle == ON {
		s.WriteString("你的小可爱提醒您:\n")
	} else {
		s.WriteString("今日分天气请查收:\n")
	}
	s.WriteString(fmt.Sprintf("当前位置:%s\n", e.SMSModel.Address))
	s.WriteString("实时天气:\n")
	s.WriteString(fmt.Sprintf("温度:%.1f°\n", e.RealtimeStruct.Result.Temperature))
	if name, ok := valueObject.RealtimeSkycon[e.RealtimeStruct.Result.Skycon]; ok {
		s.WriteString(fmt.Sprintf("天气状况:%s\n", name))
	}
	s.WriteString(fmt.Sprintf("舒适度:%s\n", e.RealtimeStruct.Result.Comfort.Desc))
	s.WriteString(fmt.Sprintf("PM25:%.1f\n", e.RealtimeStruct.Result.Pm25))
	s.WriteString("今日天气:\n")
	if len(e.DailyStruct.Result.Daily.Temperature) > 0 {
		s.WriteString(fmt.Sprintf("最低温度:%.1f°\n", e.DailyStruct.Result.Daily.Temperature[0].Min))
		s.WriteString(fmt.Sprintf("最高温度:%.1f°\n", e.DailyStruct.Result.Daily.Temperature[0].Max))
		s.WriteString(fmt.Sprintf("平均温度:%.1f°\n", e.DailyStruct.Result.Daily.Temperature[0].Avg))
	}

	if len(e.DailyStruct.Result.Daily.Skycon) > 0 {
		if name, ok := valueObject.RealtimeSkycon[e.DailyStruct.Result.Daily.Skycon[0].Value]; ok {
			s.WriteString(fmt.Sprintf("天气状况:%s\n", name))
		}
	}
	if len(e.DailyStruct.Result.Daily.Comfort) > 0 {
		s.WriteString(fmt.Sprintf("舒适度:%s\n", e.DailyStruct.Result.Daily.Comfort[0].Desc))
	}
	if len(e.DailyStruct.Result.Daily.ColdRisk) > 0 {
		s.WriteString(fmt.Sprintf("感冒指数:%s\n", e.DailyStruct.Result.Daily.ColdRisk[0].Desc))
	}
	s.WriteString("注意穿衣呦!!!")

	return s.String()
}
func (e *entity) Send(content string) {
	accountSid := viper.GetString("sms.account_sid")
	authToken := viper.GetString("sms.auth_token")
	formNumber := viper.GetString("sms.from_phone")
	urlStr := fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json", accountSid)
	msgData := url.Values{}
	fmt.Println(e.SMSModel.Phone)
	msgData.Set("To", e.SMSModel.Phone)
	msgData.Set("From", formNumber)
	msgData.Set("Body", content)
	msgDataReader := *strings.NewReader(msgData.Encode())
	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, _ := client.Do(req)
	config.Logger.Info("Send",
		zap.Any("url", urlStr),
		zap.Any("res", resp),
		zap.Any("phone", e.SMSModel.Phone),
		zap.Any("content", content),
	)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if err == nil {
			fmt.Println(data)
			fmt.Println(data["sid"])
		}
	} else {
		fmt.Printf("%+v", resp)
	}
}
func (e *entity) SmsAll(field string) ([]*model.SMSModel, error) {
	list, err := (&e.SMSModel).All(field)
	return list, err
}
func (e *entity) SetModel(model *model.SMSModel) {
	(&e.Entity).SetId(model.Id)
	e.SMSModel = *model
}
func NewEntity(id uint64) *entity {
	e := new(entity)
	(&e.Entity).SetId(id)
	return e
}

func NewEntityS(sid string) *entity {
	e := new(entity)
	(&e.Entity).SetSId(sid)
	return e
}
