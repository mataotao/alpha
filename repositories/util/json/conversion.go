package json

import jsoniter "github.com/json-iterator/go"

func Conversion(d interface{}) string {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	r, _ := json.Marshal(d)
	return string(r)
}
