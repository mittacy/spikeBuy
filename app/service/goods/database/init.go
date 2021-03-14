package database

import (
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
	"goods/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var (
	MysqlDB *gorm.DB
	RedisDB redis.Conn
)

func InitMysql() {
	if !viper.IsSet("mysql") || !viper.GetBool("mysql.enabled"){
		return
	}
	mysqlViper := viper.GetStringMapString("mysql")
	dsn := mysqlViper["user"] + ":" + mysqlViper["password"] +
		"@tcp(" + mysqlViper["host"] + ":" + mysqlViper["port"] + ")/" +
		mysqlViper["database"] + "?" + mysqlViper["config"]
	//dsn := "root:password@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	MysqlDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		logger.PanicDetail("mysql初始化失败", err)
	}
	// 设置mysql连接池
	if rawDb, err := MysqlDB.DB(); err != nil {
		logger.PanicDetail("设置Mysql连接池错误", err)
	} else {
		if exist := viper.IsSet("mysql.conn-max-idle-time"); exist {
			rawDb.SetConnMaxIdleTime(time.Second * viper.GetDuration("mysql.conn-max-idle-time")) // 连接等待时间
		}
		if exist := viper.IsSet("mysql.conn-max-life-time"); exist {
			rawDb.SetConnMaxIdleTime(time.Second * viper.GetDuration("mysql.conn-max-life-time")) // 设置了连接可复用的最大时间
		}
		if exist := viper.IsSet("mysql.max-idle-conns"); exist {
			rawDb.SetMaxIdleConns(viper.GetInt("mysql.max-idle-conns")) // 设置连接池中空闲连接的最大数量
		}
		if exist := viper.IsSet("mysql.max-open-conns"); exist {
			rawDb.SetMaxOpenConns(viper.GetInt("mysql.max-open-conns")) // 设置打开数据库连接的最大数量
		}
	}
	logger.Info("Mysql初始化成功")
}

func InitRedis() {
	if !viper.IsSet("redis") || !viper.GetBool("redis.enabled"){
		return
	}
	redisViper := viper.GetStringMapString("redis")
	var err error
	RedisDB, err = redis.Dial(redisViper["network"], redisViper["host"] + ":" + redisViper["port"])
	if err != nil {
		logger.PanicDetail("redis初始化失败", err)
	}
	logger.Info("Redis初始化成功")
}
