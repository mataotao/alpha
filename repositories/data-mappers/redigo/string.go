package redigo

import (
	redisgo "github.com/gomodule/redigo/redis"
)

/**
lockType true 加锁 false 解锁
*/
func Lock(key string, lockType bool, t int) (bool, error) {
	pool := Pool.Pool.Get()
	defer pool.Close()
	if lockType {
		lockRes, err := pool.Do("SET", key, true, "EX", t, "NX")
		if err != nil {
			return false, err
		}
		if lockRes == nil {
			return false, nil
		}
	} else {
		_, err := pool.Do("DEL", key)
		if err != nil {
			return false, err
		}
	}

	return true, nil
}
func MGet(keys []string) ([]string, error) {
	pool := Pool.Pool.Get()
	defer pool.Close()
	command := redisgo.Args{}
	for i := range keys[:] {
		command = command.AddFlat(keys[i])
	}
	res, err := redisgo.Strings(pool.Do("MGET", command...))
	if err != nil {
		return res, err
	}
	return res, nil
}
