package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"goods/app/model"
	"goods/common"
	"goods/database"
	"goods/logger"
	"time"
)

const timeLayout = "2006-01-02 15:04:05"

/**
 * @api {post} /spike 添加秒杀活动
 * @apiVersion 1.0.0
 * @apiName CreateSpike
 *
 * @apiParam {Number} goods_id 商品id
 * @apiParam {Number} price 商品价格
 * @apiParam {Number} stock 库存
 * @apiParam {Number} start_time 活动开始时间戳
 * @apiParam {Number} end_time 活动结束时间戳
 *
 * @apiParamExample {json} Request-Example:
 * {
 *     "goods_id": 1,
 *     "price": 500,
 *     "stock": 100,
 *     "start_time": 1615784400230000000,
 *     "end_time": 1615791600230000000
 * }
 *
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 OK
 * {
 *  "code": 1,
 * 	"data": null,
 * 	"msg": "成功"
 * }
 *
 */
// CreateSpike 创建秒杀活动
func CreateSpike(c *gin.Context) {
	var goods model.Spike
	msg := "成功"
	if err := c.ShouldBindJSON(&goods); err != nil {
		msg = "解析请求数据错误"
		logger.ErrorDetail(msg, err)
		common.Failed(c, -1, msg, nil)
		return
	}
	// 1. 添加到Mysql
	err := database.MysqlDB.Omit("RedisKey").Create(&goods).Error
	if err != nil {
		logger.ErrorDetail("创建sms_spike错误", err)
		return
	}
	// 2. 添加到Redis有序集合
	k := viper.GetString("redis.cacheSet")
	cacheTime := goods.StartTime - viper.GetInt64("redis.cacheAdvanceTime")
	_, err = database.RedisDB.Do("zadd", k, cacheTime, goods.Id)
	if err != nil {
		msg = "加入Redis有序集合错误"
		logger.ErrorDetail(msg, err)
		common.Failed(c, -1, msg, nil)
		return
	}
	logger.Info(fmt.Sprintf("添加商品成功, 活动开始时间: %s, 缓存库存时间: %s",
		time.Unix(goods.StartTime, 0).Format(timeLayout), time.Unix(cacheTime, 0).Format(timeLayout)))
	common.Success(c, 1, msg, nil)
}
