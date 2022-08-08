package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Iwriaw/aobi-island/gin/model"
	"github.com/Iwriaw/aobi-island/gin/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ItemRequest struct {
	Name string `query:"Name" form:"Name" json:"Name" binding:"required"`
	//稀有度 普通|精致|稀有|典藏|晨星|辉月
	Rarity string `query:"Rarity" form:"Rarity" json:"Rarity" binding:"required"`
	//分类 材料|装扮|家居|消耗
	Category string `query:"Category" form:"Category" json:"Category" binding:"required"`
	//是否占格物品
	BackpackOccupancy bool `query:"BackpackOccupancy" form:"BackpackOccupancy" json:"BackpackOccupancy" binding:"required"`
	//最大堆叠数量
	MaxStack int `query:"MaxStack" form:"MaxStack" json:"MaxStack" binding:"required"`
	//基本售价
	Price int `query:"Price" form:"Price" json:"Price" binding:"required"`
	//物品描述
	Description string `query:"Description" form:"Description" json:"Description" binding:"required"`
}

var ItemHandlerSet = HandlerSet{
	Name: "item",

	Create: func(c *gin.Context) {
		req := ItemRequest{}
		err := c.ShouldBind(&req)
		if err != nil {
			SetFailureJSONResponse(c, 400, 1, err.Error())
			return
		}
		item := model.Item{
			Name:              req.Name,
			Rarity:            req.Rarity,
			Category:          req.Category,
			BackpackOccupancy: req.BackpackOccupancy,
			MaxStack:          req.MaxStack,
			Price:             req.Price,
			Description:       req.Description,
		}
		err = service.DefaultItemService.CreateItem(&item)
		if err != nil {
			SetFailureJSONResponse(c, http.StatusInternalServerError, 1, "创建物品失败")
			return
		}
		SetSuccessJSONResponse(c, item)
	},

	Destroy: func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			SetFailureJSONResponse(c, http.StatusInternalServerError, 1, "id格式不正确")
			return
		}
		item, err := service.DefaultItemService.QueryItem(map[string]interface{}{
			"id": id,
		})
		if err == gorm.ErrRecordNotFound {
			SetFailureJSONResponse(c, http.StatusInternalServerError, 1, fmt.Sprintf("找不到id为%v的物品", id))
			return
		}
		if err != nil {
			SetFailureJSONResponse(c, http.StatusInternalServerError, 1, "查询物品失败")
			return
		}
		err = service.DefaultItemService.DeleteItem(item)
		if err != nil {
			SetFailureJSONResponse(c, http.StatusInternalServerError, 1, "删除物品失败")
			return
		}
		SetSuccessJSONResponse(c, item)
	},

	List: func(c *gin.Context) {
		items, err := service.DefaultItemService.QueryItems(map[string]interface{}{})
		if err != nil {
			SetFailureJSONResponse(c, http.StatusInternalServerError, 1, "查询物品失败")
			return
		}
		SetSuccessJSONResponse(c, items)
	},

	Retrieve: func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			SetFailureJSONResponse(c, http.StatusInternalServerError, 1, "id格式不正确")
			return
		}
		item, err := service.DefaultItemService.QueryItem(map[string]interface{}{
			"id": id,
		})
		if err == gorm.ErrRecordNotFound {
			SetFailureJSONResponse(c, http.StatusInternalServerError, 1, fmt.Sprintf("找不到id为%v的物品", id))
			return
		}
		if err != nil {
			SetFailureJSONResponse(c, http.StatusInternalServerError, 1, "查询物品失败")
			return
		}
		SetSuccessJSONResponse(c, item)
	},

	Update: func(c *gin.Context) {
		req := ItemRequest{}
		err := c.ShouldBind(&req)
		if err != nil {
			SetFailureJSONResponse(c, 400, 1, err.Error())
			return
		}
		item := model.Item{
			Name:              req.Name,
			Rarity:            req.Rarity,
			Category:          req.Category,
			BackpackOccupancy: req.BackpackOccupancy,
			MaxStack:          req.MaxStack,
			Price:             req.Price,
			Description:       req.Description,
		}
		err = service.DefaultItemService.UpdateItem(&item)
		if err != nil {
			SetFailureJSONResponse(c, http.StatusInternalServerError, 1, "创建物品失败")
			return
		}
		SetSuccessJSONResponse(c, item)
	},
}
