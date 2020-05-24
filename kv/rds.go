package kv

import (
	"errors"
	"strconv"
	"time"

	"github.com/garyburd/redigo/redis"
)

func newPoolFunc() (redis.Conn, error) {
	return redis.Dial("tcp", ":6379")
}

func CreatePool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     10,
		Dial:        newPoolFunc,
		Wait:        true,
		IdleTimeout: 240 * time.Second,
	}
}

func SetIncreasLast(rds redis.Conn, key string, val int) error {
	_, err := rds.Do("SET", key, val)
	return err
}

func GetIncreasLast(rds redis.Conn, key string) (int, error) {
	r, err := rds.Do("GET", key)
	if r == nil {
		return 0, errors.New("")
	}
	last, _ := strconv.Atoi(string(r.(interface{}).([]uint8)))
	return last, err
}

/* 查看余量 */
func SurplusMistValue(rds redis.Conn, key string) (int, error) {
	r, err := rds.Do("LLEN", key)
	if err != nil {
		return 0, err
	}
	return int(r.(int64)), nil
}

func RpopMistValue(rds redis.Conn, key string) (int, error) {
	var value int
	r, err := rds.Do("RPOP", key)
	if err != nil || r == nil {
		return value, errors.New("")
	}
	value, _ = strconv.Atoi(string(r.(interface{}).([]uint8)))
	return value, nil
}

func LpushMistValue(rds redis.Conn, key string, val int) error {
	_, err := rds.Do("LPUSH", key, val)
	return err
}
