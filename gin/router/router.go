package router

import (
	"github.com/Iwriaw/aobi-island/gin/handler"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	router := r.Group("/api/v1")
	handler.ItemHandlerSet.RegisterTo(router)
}
