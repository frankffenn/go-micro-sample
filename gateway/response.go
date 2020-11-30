package main

import (
	"github.com/frankffenn/go-micro-sample/gateway/errors"
	"github.com/gin-gonic/gin"
)

type jsonObj map[string]interface{}

func ResponseSuccess(data *jsonObj) *gin.H{
	return &gin.H{
		"success":true,
		"data": data,
	}
}

func ResponseFailWithErrorCode(code errors.ErrorCode) *gin.H {
	return  &gin.H{
		"code":code,
		"message": code.Message(),
	}
}