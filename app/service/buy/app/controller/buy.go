package controller

import (
	"buy/app/model"
	"buy/common"
	"buy/database"
	"buy/logger"
	"github.com/gin-gonic/gin"
	"time"
)

/**
 * @api {post} /spike/buy 购买秒杀商品
 * @apiVersion 1.0.0
 * @apiName Buy
 * @apiGroup Spike
 *
 * @apiParam {Number} user_id 用户id
 * @apiParam {Number} spike_id 秒杀活动id
 *
 * @apiParamExample {json} Request-Example:
 * {
 *  "user_id": 432,
 *  "spike_id": 1
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
func Buy(c *gin.Context) {
	type buyRequest struct {
		UserId  int `json:"user_id"`
		SpikeId int `json:"spike_id"`
	}
	var b buyRequest
	if err := c.ShouldBindJSON(&b); err != nil {
		common.Failed(c, -1, "参数错误", nil)
		return
	}

	/*
	 * 1. 检查时间是否处于 startTime ~ endTime
	 * 2. 减本地库存
	 * 3. 获取Redis库存队列
	 * 4. 将订单写入Kafka
	 */
	spike := model.LocalSpikeMap[b.SpikeId]
	if spike == nil || !spike.IsOnSale() {
		common.Success(c, -1, "活动未开始", nil)
		return
	}

	// 2. 减本地库存
	stock := model.LocalStockMap[b.SpikeId]
	if stock == nil || !stock.DeductionStock() {
		common.Success(c, -1, "库存不足", nil)
		return
	}

	// 3. 获取Redis库存队列
	reply, err := database.RedisPool.Get().Do("lpop", spike.RedisKey)
	if err != nil {
		common.Success(c, -1, "库存不足", nil)
		logger.ErrorDetail("获取redis库存队列错误", err)
		return
	}
	if reply == nil {
		common.Success(c, -1, "库存不足", nil)
		return
	}

	// 4. 将订单写入Kafka
	var order model.Order
	order.CreateTime = time.Now().UnixNano()
	order.UserId = b.UserId
	order.GoodsId = spike.GoodsId
	if err = database.KafkaWriteOrder(order); err != nil {
		common.Failed(c, -1, "后台错误", nil)
		logger.ErrorDetail("写入kafka订单错误", err)
		return
	}

	common.Success(c, 1, "success", nil)
}

/**
 * @api {post} /spike/cache 缓存秒杀商品库存
 * @apiVersion 1.0.0
 * @apiName CacheSpike
 * @apiGroup Spike
 *
 * @apiParam {Number} id 秒杀活动id
 * @apiParam {Number} goods_id 商品id
 * @apiParam {Number} price 商品价格
 * @apiParam {Number} stock 库存
 * @apiParam {Number} start_time 活动开始时间戳
 * @apiParam {Number} end_time 活动结束时间戳
 * @apiParam {string} redis_key Redis队列键名，代表库存
 *
 * @apiParamExample {json} Request-Example:
 * {
 *     "id": 12,
 *     "goods_id": 2,
 *     "price": 500,
 *     "stock": 100,
 *     "start_time": 1615784400230000000,
 *     "end_time": 1615791600230000000,
 *     "redis_key": "spike-stock-12"
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
func CacheSpike(c *gin.Context) {
	var spike model.Spike
	if err := c.ShouldBindJSON(&spike); err != nil {
		common.Failed(c, -1, "参数错误", nil)
		return
	}

	// 1. 加入local Stock Map
	localStock := model.NewLocalStock(spike.Stock, 0)
	model.LocalStockMap[spike.Id] = &localStock

	// 2. 加入local Info Map
	model.LocalSpikeMap[spike.Id] = &spike

	common.Success(c, 1, "缓存成功", nil)
}
