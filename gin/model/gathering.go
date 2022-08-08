package model

import "time"

// 资源采集
type Gathering struct {
	Id uint64 `gorm:"primaryKey"`
	//疲劳值增加
	FatigueIncrease int `gorm:"not null"`
	//物品Id
	ItemId uint64 `gorm:"not null"`
	//物品数量
	ItemNum int `gorm:"not null"`
	//经验奖励
	ExpReward int `gorm:"not null"`
	//每日采集次数限制
	TimesLimit int `gorm:"not null"`
	//采集次数容量
	Capacity int `gorm:"not null"`
	//次数恢复时间
	Resume time.Duration `gorm:"not null"`
}
