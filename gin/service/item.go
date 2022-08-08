package service

import (
	"github.com/Iwriaw/aobi-island/gin/dao"
	"github.com/Iwriaw/aobi-island/gin/model"
)

type ItemService struct {
	ItemDAO *dao.ItemDAO
}

func (i *ItemService) CreateItem(item *model.Item) error {
	err := i.ItemDAO.CreateItem(item)
	if err != nil {
		return err
	}
	return nil
}

func (i *ItemService) DeleteItem(item *model.Item) error {
	err := i.ItemDAO.DeleteItem(item)
	if err != nil {
		return err
	}
	return nil
}

func (i *ItemService) UpdateItem(item *model.Item) error {
	err := i.ItemDAO.UpdateItem(item)
	if err != nil {
		return err
	}
	return nil
}

func (i *ItemService) QueryItem(query map[string]interface{}) (*model.Item, error) {
	var item *model.Item
	item, err := i.ItemDAO.QueryItem(query)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (i *ItemService) QueryItems(query map[string]interface{}) (*[]model.Item, error) {
	var items *[]model.Item
	items, err := i.ItemDAO.QueryItems(query)
	if err != nil {
		return nil, err
	}
	return items, nil
}
