package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"goods/app/model"
	"goods/common"
	"goods/database"
	"goods/logger"
)

/**
 * @api {post} /sms_goods 添加秒杀活动
 * @apiVersion 1.0.0
 * @apiName CreateSmsGoods
 * @apiGroup SmsGoods
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
// CreateSmsGoods 创建秒杀活动接口
func CreateSmsGoods(c *gin.Context) {
	var goods model.SmsGoods
	msg := "成功"
	if err := c.ShouldBindJSON(&goods); err != nil {
		msg = "解析请求数据错误"
		logger.ErrorDetail(msg, err)
		common.Failed(c, -1, msg, nil)
		return
	}
	// 1. 添加到Mysql
	err := database.MysqlDB.Create(&goods).Error
	if err != nil {
		logger.ErrorDetail("创建sms_goods错误", err)
		return
	}
	// 2. 添加到Redis有序集合
	k := viper.GetString("redis.cacheSet")
	advance := viper.GetInt64("redis.cacheAdvanceTime")
	cacheTime := goods.StartTime-advance*1000000000
	_, err = database.RedisDB.Do("zadd", k, cacheTime, goods.Id)
	if err != nil {
		msg = "加入Redis有序集合错误"
		logger.ErrorDetail(msg, err)
		common.Failed(c, -1, msg, nil)
		return
	}
	common.Success(c, 1, msg, nil)
}
