package api

import (
	"github.com/gin-gonic/gin"
	"github.com/icepie/lit-edu-go/model"
	"github.com/icepie/lit-edu-go/service"
)

// JWGetStatus 获取教务服务状态
func JWGetStatus(c *gin.Context) {
	var service service.GetStatusService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetStatus()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// JWGetBaseInfo 通过学号获取学生基本信息
func JWGetBaseInfo(c *gin.Context) {
	var service service.GetBaseInfoService

	// 传入用户
	user, _ := c.Get("user")
	service.User = user.(model.StuAccount)

	if err := c.ShouldBind(&service); err == nil {
		res := service.GetBaseInfo()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// JWGetScore 通过学号获取成绩
func JWGetScore(c *gin.Context) {
	var service service.GetScoreService

	// 传入用户
	user, _ := c.Get("user")
	service.User = user.(model.StuAccount)

	if err := c.ShouldBind(&service); err == nil {
		res := service.GetScore()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
