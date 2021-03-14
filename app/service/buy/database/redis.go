package database

import (
	"buy/logger"
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
	"time"
)

var (
	RedisPool *redis.Pool
)

func InitRedis() {
	if err := initRedisPool(); err != nil {
		logger.PanicDetail("初始化redis错误", err)
	}
	logger.Info("redis连接池初始化成功")
}

func initRedisPool() (err error){
	redisViper := viper.Sub("redis")
	RedisPool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			var c redis.Conn
			c, err = redis.Dial(redisViper.GetString("network"),
				redisViper.GetString("host") + ":" + redisViper.GetString("port"))
			if err != nil {
				return nil, err
			}
			//if _, err := c.Do("AUTH", redisViper.GetString("pass")); err != nil {
			//	c.Close()
			//	return nil, err
			//}
			return c, nil
		},
		MaxIdle:         redisViper.GetInt("max-idle"),
		MaxActive:       redisViper.GetInt("max-active"),
		IdleTimeout:     time.Second * redisViper.GetDuration("idle-timeout"),
		Wait:            true,
	}
	if _, err := RedisPool.Get().Do("ping"); err != nil {
		logger.PanicDetail("redis连接池无法正常工作", err)
	}
	return
}
