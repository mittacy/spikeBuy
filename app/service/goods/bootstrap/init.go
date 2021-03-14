package bootstrap

import (
	"fmt"
	"goods/config"
	"goods/database"
	"goods/logger"
)

func Init() {
	fmt.Println("初始化工作...")
	// 1. 初始化配置文件
	config.InitConfig()
	// 2. 初始化日志
	logger.InitLogger()
	// 3. 初始化数据库连接 mysql、redis
	database.InitMysql()
	database.InitRedis()
	fmt.Println("初始化工作完成")
}
