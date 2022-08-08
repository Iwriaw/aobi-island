package model

type Item struct {
	Id uint64 `gorm:"primaryKey"`
	//物品名称
	Name string `gorm:"unique,size:64,not null"`
	//稀有度 普通|精致|稀有|典藏|晨星|辉月
	Rarity string `gorm:"size:64,not null"`
	//分类 材料|装扮|家居|消耗
	Category string `gorm:"size:32,not null"`
	//是否占格物品
	BackpackOccupancy bool `gorm:"not null"`
	//最大堆叠数量
	MaxStack int `gorm:"not null"`
	//基本售价
	Price int `gorm:"not null"`
	//物品描述
	Description string `gorm:"not null"`
}
