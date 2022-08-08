package dao

import (
	"github.com/Iwriaw/aobi-island/gin/config"
)

var DefaultItemDAO = ItemDAO{
	DB: config.DB,
}

func init() {
}
