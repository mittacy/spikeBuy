package cron

import (
	"context"
	"encoding/json"
	"fmt"
	"goods/app/model"
	"goods/database"
	"goods/logger"
	"time"
)

/*
 * 消费Kafka的订单数据，记录到Mysql中
 */
func StartOrderConsumer(groupId string) {
	r := database.NewKafkaReader(groupId)
	logger.Info("启动kafka订单消费线程, groupId: " + groupId)
	ctx := context.Background()
	for {
		/*
		 * 1. 阻塞获取订单消息
		 * 2. 处理订单到Mysql
		 * 3. 提交到Kafka确认消费完成
		 */
		msg, err := r.FetchMessage(ctx)
		if err != nil {
			logger.PanicDetail("获取kafka订单消息错误", err)
			time.Sleep(time.Minute)
			continue
		}

		var order model.Order
		if err = json.Unmarshal(msg.Value, &order); err != nil {
			logger.ErrorDetail("处理Kafka订单错误", err)
			return
		}
		fmt.Println("订单: ", order)
		// todo 实务操作: 减少秒杀活动的库存，添加订单到Mysql



		if err = r.CommitMessages(ctx, msg); err != nil {
			logger.PanicDetail("确认消费完Kafka订单消息失败", err)
			time.Sleep(time.Minute)
			continue
		}
	}
}
