package server

import (
	"github.com/icepie/lit-edu-go/api"

	"github.com/gin-gonic/gin"
)

// NewRouter 创建新路由
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 路由
	v1 := r.Group("api/v1")
	{
		// 测试连接
		v1.GET("ping", api.PingPong)
		jw := v1.Group("jw")
		{
			// 测试教务服务连接性
			jw.GET("status", api.JWIsWork)
			// 获取学生基础个人信息
			jw.POST("profile", api.JWGetBaseInfo)
			// 获取学生成绩
			jw.POST("score", api.JWGetScore)
		}

	}

	return r
}
