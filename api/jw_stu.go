package api

import (
	"github.com/gin-gonic/gin"
	"github.com/icepie/lit-edu-go/model"
	"github.com/icepie/lit-edu-go/service"
)

// JWUserAuth 通用认证接口
func JWUserAuth(c *gin.Context) {
	var service service.UserAuthService
	if err := c.ShouldBind(&service); err == nil {
		res := service.UserAuth()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// JWAISchedule 获取小爱课程表格式的数据
func JWAISchedule(c *gin.Context) {
	var service service.GetAIscheduleService

	// 传入用户
	user, _ := c.Get("user")
	service.User = user.(model.StuAccount)

	if err := c.ShouldBind(&service); err == nil {
		res := service.GetAIschedule()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
