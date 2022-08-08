package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseJSONResponse struct {
	Code int
	Msg  string
	Data interface{}
}

func SetSuccessJSONResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, BaseJSONResponse{
		Code: 0,
		Msg:  "success",
		Data: data,
	})
}

func SetFailureJSONResponse(c *gin.Context, httpStatus int, code int, msg string) {
	c.JSON(httpStatus, BaseJSONResponse{
		Code: code,
		Msg:  msg,
	})
}
