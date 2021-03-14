package model

type Spike struct {
	Id         int32 `json:"id" gorm:"primaryKey"`
	GoodsId    int64 `json:"goods_id"`
	Price      int64 `json:"price"`
	Stock      int32 `json:"stock"`
	StartTime  int64 `json:"start_time"`
	EndTime    int64 `json:"end_time"`
}

func (*Spike) TableName() string {
	return "sms_spike"
}
