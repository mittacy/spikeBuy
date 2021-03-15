package cron

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
	"goods/app/model"
	"goods/common"
	"goods/database"
	"goods/logger"
	"strconv"
	"time"
)

func InitCacheStock() {
	// 检查Redis有序集合缓存时间是否到点了，到点了从有序集合取出进行操作并从redis集合中删除
	redisViper := viper.Sub("redis")
	for {
		t := time.Now().Unix()
		spikesId, err := redis.Ints(database.RedisDB.Do("zrangebyscore", redisViper.GetString("cacheSet"), 0, t))
		if err != nil {
			logger.ErrorDetail("redis查询库存有序集合错误", err)
			time.Sleep(redisViper.GetDuration("cacheAdvanceTime") / 2 * time.Second)
			continue
		}
		if len(spikesId) > 0 {
			// 启动线程去处理每个到点的缓存
			for _, v := range spikesId {
				go handleCacheStock(v)
			}
			// 删除到点缓存
			_, err = database.RedisDB.Do("zrem", redis.Args{}.Add(redisViper.GetString("cacheSet")).AddFlat(spikesId)...)
			if err != nil {
				logger.ErrorDetail("redis删除有序集合错误", err)
			}
		}
		// 每次监听间隔时间，要保证不大于配置文件的开售提前缓存时间
		var sleepTime = 10 * time.Second
		if sleepTime >= redisViper.GetDuration("cacheAdvanceTime") {
			sleepTime = redisViper.GetDuration("cacheAdvanceTime") / 2
		}
		time.Sleep(time.Second * sleepTime)
	}
}

func handleCacheStock(id int) {
	/*
	 * 1. 根据id请求Mysql获取秒杀商品信息
	 * 2. 向 服务发现系统 获取抢购服务数
	 * 3. 计算每个服务应该缓存的库存
	 * 4. 缓存到Redis队列代表库存
	 * 5. 调用各个服务推送缓存
	 */

	// 1. 根据id请求Mysql获取秒杀商品信息
	var spike model.Spike
	spike.Id = id
	if err := database.MysqlDB.First(&spike).Error; err != nil {
		logger.ErrorDetail("查询秒杀活动错误", err)
		return
	}
	// 2. 向 服务发现系统 获取抢购服务数
	services := database.GetAllBuyInstance()
	// 3. 计算每个服务应该缓存的库存
	n := len(services)
	if n == 0 {
		return
	}
	panicCount := viper.GetInt("discovery.instance.panic-count")
	if panicCount <= 0 {
		logger.Error("没有抢购服务")
		return
	}
	oneServiceStock := (1 + panicCount) * spike.Stock / n
	logger.Info(
		fmt.Sprintf("Id为 %d 的活动需要进行缓存了, 该商品总库存: %d; " +
			"当前抢购服务数: %d, 容灾宕机数: %d, 每台服务能够发起 %d 次redis减库请求",
			id, spike.Stock, n, panicCount, oneServiceStock))
	// 4. 缓存到Redis队列代表库存
	stocks := make([]int, spike.Stock)
	logger.Info("总库存: " + strconv.Itoa(len(stocks)))
	spike.Stock = oneServiceStock
	spike.RedisKey = "spike-buy-" + strconv.Itoa(id)
	_, err := database.RedisDB.Do("LPUSH", redis.Args{spike.RedisKey}.AddFlat(stocks)...)
	if err != nil {
		logger.ErrorDetail("推送库存队列到Redis失败", err)
		return
	}
	// 5. 调用各个服务推送缓存
	spikeJson, _ := json.Marshal(&spike)
	for _, v := range services {
		go handle(v, id, spikeJson)
	}
}

func handle(host string, id int, data []byte) {
	url := "http://" + host + viper.GetString("discovery.instance.cache-api")
	result, err := common.PostRequest(url, data)
	if err != nil {
		logger.ErrorDetail(fmt.Sprintf("推送活动Id: %d 的缓存失败, 服务IP: %s,", id, host), err)
	}
	if result.Code < 0 {
		logger.Error(fmt.Sprintf("推送缓存, 服务IP: %s, 失败", host))
	} else {
		logger.Info(fmt.Sprintf("推送活动Id: %d 的缓存成功, 服务IP: %s", id, host))
	}
}
