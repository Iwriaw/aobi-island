package model

//制造配方
type RecipeRelation struct {
	Id uint64 `gorm:"primaryKey"`
	//输入物品Id
	InputItemId int `gorm:"not null"`
	//输入物品数量
	InputItemNum int `gorm:"not null"`
	//输出物品Id
	OutputItemId int `gorm:"not null"`
	//输出物品数量
	OutputItemNum int `gorm:"not null"`
	//生产者
	Producer string `gorm:"size:32,not null"`
	//金币花费
	GoldCost int `gorm:"not null"`
	//经验奖励
	ExpReward int `gorm:"not null"`
}
