package bootstrap

import (
	"fmt"
	"github.com/spf13/viper"
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
	// 4. 初始化服务发现
	database.InitNacos()

	time.Sleep(time.Second)
	// 5. 启动后台监听、处理线程
	groupId := "mysql-spike-buy"
	n := viper.GetInt("order.thread")
	for i := 0; i < n; i++ {
		go cron.StartOrderConsumer(groupId)
	}
	// 6. 监听Redis有序集合,处理分发库存
	go cron.InitCacheStock()
	logger.Info("初始化工作完成")
}
