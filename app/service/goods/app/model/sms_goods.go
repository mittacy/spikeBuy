package model

import "time"

type SmsGoods struct {
	Id         int32 `json:"id" gorm:"primaryKey"`
	CreateTime int64 `json:"create_time" gorm:"autoCreateTime"`
	UpdateTime int64 `json:"update_time" gorm:"autoUpdateTime"`
	GoodsId    int64 `json:"goods_id"`
	Price      int64 `json:"price"`
	Stock      int32 `json:"stock"`
	StartTime  int64 `json:"start_time"`
	EndTime    int64 `json:"end_time"`
}

func (*SmsGoods) TableName() string {
	return "sms_goods"
}

func (sg *SmsGoods) IsOnSale() bool {
	t := time.Now().Unix()
	return t >= sg.StartTime && t < sg.EndTime
}
