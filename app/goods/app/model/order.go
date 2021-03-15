package model

type Order struct {
	SpikeId    int   `json:"spike_id"`
	CreateTime int64 `json:"create_time"`
	UserId     int   `json:"user_id"`
	GoodsId    int64 `json:"goods_id"`
}

func (*Order) TableName() string {
	return "oms_order"
}
