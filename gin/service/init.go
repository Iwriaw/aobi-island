package service

import "github.com/Iwriaw/aobi-island/gin/dao"

var DefaultItemService = ItemService{
	ItemDAO: &dao.DefaultItemDAO,
}

func init() {

}
