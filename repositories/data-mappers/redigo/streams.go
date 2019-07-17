package redigo

import (
	redisgo "github.com/gomodule/redigo/redis"
)

const (
	QueueTimeOut int64 = 30000
)

type StreamValue struct {
	Field string
	Value string
}

type StreamData struct {
	Id string `json:"id"`
	StreamValue
}

func XAdd(key, id string, value ...StreamValue) error {
	pool := Pool.Pool.Get()
	defer pool.Close()

	command := redisgo.Args{}.Add(key).AddFlat(id)

	for i := range value[:] {
		command = command.AddFlat(value[i].Field)
		command = command.AddFlat(value[i].Value)
	}

	if _, err := pool.Do("XADD", command...); err != nil {
		return err
	}
	return nil
}

func XGroup(action, streams, name, id string) error {
	pool := Pool.Pool.Get()
	defer pool.Close()
	if _, err := pool.Do("XGROUP", action, streams, name, id); err != nil {
		return err
	}
	return nil
}

func XReadGroupAck(group, consumer, id string, key []string, count, block int64) (map[string][]map[string]string, error) {
	data := make(map[string][]map[string]string)
	pool := Pool.Pool.Get()
	defer pool.Close()
	command := redisgo.Args{}.Add("GROUP").AddFlat(group)
	command = command.AddFlat(consumer)
	command = command.AddFlat("COUNT").AddFlat(count)
	command = command.AddFlat("BLOCK").AddFlat(block)
	command = command.AddFlat("STREAMS")
	for i := range key[:] {
		command = command.AddFlat(key[i])
	}
	command = command.AddFlat(id)
	d, err := redisgo.Values(pool.Do("XREADGROUP", command...))
	if err != nil {
		return data, err
	}
	if len(d) == 0 {
		return data, nil
	}
	pool.Send("MULTI")
	for i := range d[:] {
		keyGroup, err := redisgo.Values(d[i], nil)
		if err != nil {
			return data, err
		}
		k, err := redisgo.String(keyGroup[0], nil)
		if err != nil {
			return data, err
		}
		keyValues, err := redisgo.Values(keyGroup[1], nil)
		if err != nil {
			return data, err
		}
		for v := range keyValues[:] {
			idLevel, err := redisgo.Values(keyValues[v], nil)
			if err != nil {
				return data, err
			}
			valueId, err := redisgo.String(idLevel[0], nil)
			if err != nil {
				return data, err
			}
			valueDatas, err := redisgo.Values(idLevel[1], nil)
			if err != nil {
				return data, err
			}
			streamData := make(map[string]string)
			streamData["id"] = valueId
			pool.Send("XACK", k, group, valueId)
			for j := range valueDatas[:] {
				if j&1 == 1 {
					continue
				}
				field, err := redisgo.String(valueDatas[j], nil)
				if err != nil {
					return data, err
				}
				fValue, err := redisgo.String(valueDatas[j+1], nil)
				if err != nil {
					return data, err
				}
				streamData[field] = fValue
			}
			data[k] = append(data[k], streamData)
		}

	}
	pool.Do("EXEC")
	if err != nil {
		return data, err
	}
	return data, nil
}
