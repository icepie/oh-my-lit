package api

import (
	"github.com/gin-gonic/gin"
	"github.com/icepie/lit-edu-go/model/serializer"
	"github.com/icepie/lit-edu-go/service/jw"
)

// JWIsWork 测试教务服务是否正常
func JWIsWork(c *gin.Context) {
	if iswork, err := jw.IsWork(); iswork && err == nil {
		c.JSON(200, serializer.Response{
			Status: 200,
			Msg:    "LIT JW is work fine!",
		})
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}
