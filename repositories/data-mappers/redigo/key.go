package redigo

import (
	"alpha/repositories/util/slice"
	redisgo "github.com/gomodule/redigo/redis"
)

func Exists(key string) (bool, error) {
	pool := Pool.Pool.Get()
	defer pool.Close()
	result, err := redisgo.Bool(pool.Do("EXISTS", key))
	return result, err
}
func Scan(pattern string, count int) ([]string, error) {
	pool := Pool.Pool.Get()
	defer pool.Close()
	start := 1
	index := 0
	data := make([]string, 0)
	for start == 1 {
		res, err := redisgo.Values(pool.Do("SCAN", redisgo.Args{}.Add(index).AddFlat("match").AddFlat(pattern).AddFlat("count").AddFlat(count)...))
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
func Scans(patterns []string, count int) ([]string, error) {
	pool := Pool.Pool.Get()
	defer pool.Close()
	data := make([]string, 0)
	for i := range patterns[:] {
		start := 1
		index := 0
		for start == 1 {
			res, err := redisgo.Values(pool.Do("SCAN", redisgo.Args{}.Add(index).AddFlat("match").AddFlat(patterns[i]).AddFlat("count").AddFlat(count)...))
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
	}

	//去重
	res := slice.RemoveDuplicateElementString(data)
	return res, nil
}
