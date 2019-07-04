package redigo

import (
	redisgo "github.com/gomodule/redigo/redis"
)

func Lpush(key string, value ...string) error {
	pool := Pool.Pool.Get()
	defer pool.Close()
	command := redisgo.Args{}.Add(key)

	for i := range value[:] {
		command = command.AddFlat(value[i])
	}

	_, err := pool.Do("LPUSH", command...)
	if err != nil {
		return err
	}
	return nil
}

func Rpop(key string) (string, error) {
	pool := Pool.Pool.Get()
	defer pool.Close()
	value, err := pool.Do("RPOP", key)
	if err != nil {
		return "", err
	}
	if value == nil {
		return "", nil
	}
	res, err := redisgo.String(value, err)
	if err != nil {
		return "", err
	}
	return res, nil
}
