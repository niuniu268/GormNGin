package controllers

import (
	"github.com/gin-gonic/gin"
)

type JsonStruct struct {
	Code  int         `json:"code"`
	Msg   interface{} `json:"msg"`
	Data  interface{} `json:"data"`
	Count int         `json:"count"`
}

func ReturnSuccess(context *gin.Context, code int, msg interface{}, data interface{}, count int) {

	json := &JsonStruct{
		Code:  code,
		Msg:   msg,
		Data:  data,
		Count: count,
	}

	context.JSON(200, json)
}

func ReturnError(context *gin.Context, code int, msg interface{}) {
	json := &JsonStruct{
		Code: code,
		Msg:  msg,
	}

	context.JSON(200, json)
}
