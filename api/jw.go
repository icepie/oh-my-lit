package api

import (
	"github.com/gin-gonic/gin"
	"github.com/icepie/lit-edu-go/model"
	"github.com/icepie/lit-edu-go/service"
	"github.com/icepie/lit-edu-go/service/jw"
)

// JWIsWork 测试教务服务是否正常
func JWIsWork(c *gin.Context) {
	if iswork, err := jw.IsWork(); iswork && err == nil {
		c.JSON(200, model.Response{
			Status: 200,
			Data:   "",
			Msg:    "lit jw is work fine!",
		})
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}

// JWGetScore 通过学号获取成绩
func JWGetScore(c *gin.Context) {
	var service service.GetScoreService
	// 获取token
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetScore()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
