package handler

import (
	"github.com/gin-gonic/gin"
)

type HandlerSet struct {
	Name     string
	Create   func(c *gin.Context)
	Destroy  func(c *gin.Context)
	List     func(c *gin.Context)
	Retrieve func(c *gin.Context)
	Update   func(c *gin.Context)
}

func (h *HandlerSet) GetName() string {
	return h.Name
}

func (h *HandlerSet) RegisterTo(r interface{}) {
	var router *gin.RouterGroup
	switch r.(type) {
	case *gin.Engine:
		router = r.(*gin.Engine).Group(h.Name)
	case *gin.RouterGroup:
		router = r.(*gin.RouterGroup).Group(h.Name)
	}
	router.POST("/", h.Create)
	router.DELETE("/:id", h.Destroy)
	router.GET("/", h.List)
	router.GET("/:id", h.Retrieve)
	router.PUT("/:id", h.Update)
}
