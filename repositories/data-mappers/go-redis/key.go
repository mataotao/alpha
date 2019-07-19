package redigo

import (
	"alpha/repositories/util/slice"
)

func Scan(match string, count int64) ([]string, error) {
	var cursor uint64
	data := make([]string, 0)
	for {
		keys, index, err := Client.Client.Scan(cursor, match, count).Result()
		if err != nil {
			return data, err
		}
		cursor = index
		for i := range keys[:] {
			data = append(data, keys[i])
		}
		if index == 0 {
			break
		}
	}
	//去重
	res := slice.RemoveDuplicateElementString(data)
	return res, nil
}
