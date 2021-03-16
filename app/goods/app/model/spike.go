package model

type Spike struct {
	Id        int    `json:"id" gorm:"primaryKey"`
	GoodsId   int64  `json:"goods_id"`
	Price     int64  `json:"price"`
	Stock     int    `json:"stock"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
	RedisKey  string `json:"redis_key"`
}

func (*Spike) TableName() string {
	return "sms_spike"
}
