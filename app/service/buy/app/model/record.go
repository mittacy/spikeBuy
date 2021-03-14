package model

import (
	"buy/logger"
)

var (
	// LocalStockMap 秒杀商品库存Map
	// key: int SpikeId 秒杀活动id
	// value: LocalStock 秒杀商品库存结构体
	LocalStockMap map[int]*LocalStock
	// LocalSpikeMap 秒杀商品信息Map
	// key: int SpikeId 秒杀活动id
	// value: Spike 秒杀商品信息结构体
	LocalSpikeMap map[int]*Spike
)


func InitLocalGoodsStock() {
	LocalStockMap = make(map[int]*LocalStock)
	LocalSpikeMap = make(map[int]*Spike)
	logger.Info("初始化秒杀商品本地缓存成功")
}
