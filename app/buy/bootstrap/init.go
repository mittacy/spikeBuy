package bootstrap

import (
	"buy/app/cron"
	"buy/app/model"
	"fmt"
	"buy/config"
	"buy/database"
	"buy/logger"
)

func Init() {
	fmt.Println("初始化工作...")
	// 1. 初始化配置文件
	config.InitConfig()
	// 2. 初始化日志
	logger.InitLogger()
	// 3. 初始化数据库连接redis
	database.InitRedis()
	// 4. 初始化本地库存
	model.InitLocalGoodsStock()
	// 5. 初始化Kafka
	database.InitKafka()
	// 6. 服务注册
	cron.InitNacos()
	logger.Info("初始化工作完成")
}
