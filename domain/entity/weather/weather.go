package weather

import (
	baseEntity "alpha/domain/entity"

	"github.com/spf13/viper"

	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type entity struct {
	baseEntity.Entity
}

func (e *entity) Content() string {
	s := "你的小可爱提醒您:\n" +
		"当前位置:绿地468\n" +
		"实时天气:\n" +
		"温度:15°\n" +
		"气象:多云（夜间）\n" +
		"舒适度:凉爽\n" +
		"PM25:41\n" +
		"今日天气:\n" +
		"最低温度:17°\n" +
		"最高温度:23.16°\n" +
		"平均温度:19.15°\n" +
		"天气状况:多云（白天）\n" +
		"舒适度:温暖\n" +
		"感冒指数:易发\n" +
		"注意穿衣呦！！！\n"
	return s
}
func (e *entity) Send(phone, content string) {
	accountSid := viper.GetString("sms.account_sid")
	authToken := viper.GetString("sms.auth_token")
	formNumber := viper.GetString("sms.from_phone")
	urlStr := fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json", accountSid)
	msgData := url.Values{}
	msgData.Set("To", "+8617628293814")
	msgData.Set("From", formNumber)
	msgData.Set("Body", content)
	msgDataReader := *strings.NewReader(msgData.Encode())
	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, _ := client.Do(req)
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
