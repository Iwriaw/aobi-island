package dao

import (
	"fmt"

	"github.com/Iwriaw/aobi-island/gin/model"
	"gorm.io/gorm"
)

type ItemDAO struct {
	DB *gorm.DB
}

func (i *ItemDAO) CreateItem(item *model.Item) error {
	err := i.DB.Create(item).Error
	return err
}

func (i *ItemDAO) CreateItems(items *[]model.Item) error {
	err := i.DB.Create(items).Error
	return err
}

func (i *ItemDAO) UpdateItem(item *model.Item) error {
	err := i.DB.Save(item).Error
	return err
}

func (i *ItemDAO) UpdateItems(items *[]model.Item) error {
	err := i.DB.Save(items).Error
	return err
}

func (i *ItemDAO) DeleteItem(item *model.Item) error {
	err := i.DB.Delete(item).Error
	return err
}

func (i *ItemDAO) DeleteItems(items *[]model.Item) error {
	err := i.DB.Delete(items).Error
	return err
}

func (i *ItemDAO) QueryItem(query map[string]interface{}) (*model.Item, error) {
	item := model.Item{}
	err := i.DB.Where(query).Take(&item).Error
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &item, nil
}

func (i *ItemDAO) QueryItems(query map[string]interface{}) (*[]model.Item, error) {
	items := make([]model.Item, 0)
	err := i.DB.Where(query).Find(&items).Error
	if err != nil {
		return nil, err
	}
	return &items, nil
}
