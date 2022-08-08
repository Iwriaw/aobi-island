package model

//订单
type Order struct {
	Id uint64 `gorm:"primaryKey"`
	//物品Id
	ItemId uint64
	//金币奖励
	GoldReward int
	//经验奖励
	ExpReward int
}
