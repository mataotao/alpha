package redigo

import (
	"alpha/pkg/constvar"
	redisgo "github.com/gomodule/redigo/redis"
)

//按照分值从大到小
func ZRevRange(key string, page, limit uint64) ([]string, error) {
	pool := Pool.Pool.Get()
	defer pool.Close()

	if page == 0 {
		page = 1
	}
	if limit == 0 || limit > constvar.MaxLimit {
		limit = constvar.DefaultLimit
	}
	start := (page - 1) * limit
	stop := start + limit - 1
	command := redisgo.Args{}.Add(key).AddFlat(start).AddFlat(stop)
	result, err := redisgo.Strings(pool.Do("ZREVRANGE", command...))
	return result, err
}
func ZRevRangeAll(key string, start, stop int) ([]string, error) {
	pool := Pool.Pool.Get()
	defer pool.Close()
	command := redisgo.Args{}.Add(key).AddFlat(start).AddFlat(stop)
	result, err := redisgo.Strings(pool.Do("ZREVRANGE", command...))
	return result, err
}

//按照分值从小到大
func ZRange(key string, page, limit uint64) ([]string, error) {
	pool := Pool.Pool.Get()
	defer pool.Close()

	if page == 0 {
		page = 1
	}
	if limit == 0 || limit > constvar.MaxLimit {
		limit = constvar.DefaultLimit
	}
	start := (page - 1) * limit
	stop := start + limit - 1
	command := redisgo.Args{}.Add(key).AddFlat(start).AddFlat(stop)
	result, err := redisgo.Strings(pool.Do("ZRANGE", command...))
	return result, err
}

//获取key对应的值
func ZRevRangeKeys(keys []string, start, end int) ([]string, error) {
	data := make([]string, 0)
	pool := Pool.Pool.Get()
	defer pool.Close()
	pool.Send("MULTI")
	for i := range keys[:] {
		pool.Send("ZREVRANGE", keys[i], start, end)
	}
	res, err := redisgo.Values(pool.Do("EXEC"))
	if err != nil {
		return data, err
	}
	for i := range res[:] {
		values, err := redisgo.Strings(res[i], err)
		if err != nil {
			return data, err
		}
		for j := range values[:] {
			data = append(data, values[j])
		}

	}
	return data, nil
}
