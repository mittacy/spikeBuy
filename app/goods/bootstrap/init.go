package bootstrap

import (
	"fmt"
	"goods/app/cron"
	"goods/config"
	"goods/database"
	"goods/logger"
	"time"
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
	// 4. 监听Redis有序集合,处理分发库存
	database.InitNacos()
	time.Sleep(time.Second)
	cron.InitCacheStock()
	logger.Info("初始化工作完成")
}
