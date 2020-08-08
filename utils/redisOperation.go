package utils

//redis操作工具类
import (
	"MarvelousBlog-Backend/common"
	"MarvelousBlog-Backend/config"
	"github.com/garyburd/redigo/redis"
)

//封装的Redis Get操作工具函数
func RedisGet(key string) (string, error) {
	conn := config.RedisPool.Get()
	defer conn.Close()
	value, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return "", err
	}
	return value, err
}

//封装的Redis Set操作工具函数
func RedisSet(key string, value interface{}) error {
	conn := config.RedisPool.Get()
	value, _ = config.Json.Marshal(value)
	defer conn.Close()
	_, err := conn.Do("SET", key, value)
	if err != nil {
		return err
	}
	//设置redis过期时间
	_, err = conn.Do("EXPIRE", key, common.REDIS_EXPIRE_TIME)
	return err
}

func RedisDelete(key string) error {
	conn := config.RedisPool.Get()
	defer conn.Close()
	_, err := conn.Do("DEL", key)
	return err
}
