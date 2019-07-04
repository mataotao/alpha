package redigo

import (
	"alpha/repositories/util/slice"
	redisgo "github.com/gomodule/redigo/redis"
)

func Sscan(key, pattern string, count int) ([]string, error) {
	pool := Pool.Pool.Get()
	defer pool.Close()
	start := 1
	index := 0
	data := make([]string, 0)
	for start == 1 {
		res, err := redisgo.Values(pool.Do("SSCAN", redisgo.Args{}.Add(key).AddFlat(index).AddFlat("match").AddFlat(pattern).AddFlat("count").AddFlat(count)...))
		if err != nil {
			return data, err
		}
		r, err := redisgo.Strings(res[1], err)
		if err != nil {
			return data, err
		}
		promotionOnceIndex, err := redisgo.Int(res[0], err)
		if err != nil {
			return data, err
		}
		index = promotionOnceIndex
		if index == 0 {
			start = 0
		}
		for i := range r[:] {
			data = append(data, r[i])
		}
	}
	//去重
	res := slice.RemoveDuplicateElementString(data)
	return res, nil

}

func Srandmember(key string, count int64) ([]string, error) {
	pool := Pool.Pool.Get()
	defer pool.Close()
	data, err := redisgo.ByteSlices(pool.Do("SRANDMEMBER", key, count))
	if err != nil {
		return []string{}, err
	}
	res := make([]string, len(data))
	for i := range data[:] {
		res[i] = string(data[i])
	}
	return res, err
}
