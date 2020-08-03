package utils

//redis操作工具类
import (
	"MarvelousBlog-Backend/config"
	"github.com/garyburd/redigo/redis"
)

//封装的Redis Get操作工具函数
func RedisGet(key string) string {
	conn := config.RedisPool.Get()
	defer conn.Close()
	value, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return ""
	}
	return value
}

//封装的Redis Set操作工具函数
func RedisSet(key string, value string) error {
	conn := config.RedisPool.Get()
	defer conn.Close()
	_, err := conn.Do("SET", key, value)
	if err != nil {
		return err
	}
	return nil
}
