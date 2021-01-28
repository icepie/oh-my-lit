package api

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/icepie/lit-edu-go/model"
	"github.com/icepie/lit-edu-go/pkg/e"
	"gopkg.in/go-playground/validator.v8"
)

const (
	// Version of lit-edu-go
	Version = "v0.1.7"
)

// PingPong 测试连接接口
func PingPong(c *gin.Context) {
	c.JSON(200, model.Response{
		Status: 200,
		Data:   "",
		Msg:    "pong",
	})
}

// ErrorResponse 返回错误消息
func ErrorResponse(err error) model.Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, err := range ve {
			code := e.INVALID_PARAMS
			return model.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  fmt.Sprint(err),
			}
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		code := e.INVALID_PARAMS
		return model.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  fmt.Sprint(err),
		}
	}

	code := e.INVALID_PARAMS
	return model.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Error:  fmt.Sprint(err),
	}
}
