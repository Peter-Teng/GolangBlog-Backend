package config

import (
	"github.com/garyburd/redigo/redis"
)

var RedisPool *redis.Pool

//初始化一个redis连接池
func init() {
	redisCon := RedisHost + ":" + RedisPort
	RedisPool = &redis.Pool{
		MaxIdle:     8,   //最大空闲数
		MaxActive:   0,   //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: 300, //连接关闭时间 300秒 （300秒不使用自动关闭）
		//开启redis连接 & 使用密码
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", redisCon)
			if err != nil {
				return nil, err
			}
			if RedisPassword != "" {
				if _, err := conn.Do("AUTH", RedisPassword); err != nil {
					conn.Close()
					return nil, err
				}
			}
			return conn, err
		},
	}
}
