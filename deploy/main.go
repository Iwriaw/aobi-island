package main

import (
	"fmt"

	"github.com/Iwriaw/aobi-island/gin/config"
	"github.com/Iwriaw/aobi-island/gin/model"
)

func main() {
	err := config.DB.AutoMigrate(&model.Item{})
	if err != nil {
		fmt.Println(err)
	}
}
