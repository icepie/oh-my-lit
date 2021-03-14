package server

import (
	"github.com/icepie/lit-edu-go/api"

	"github.com/gin-gonic/gin"
)

// ApiRouter 创建新路由
func ApiRouter() *gin.Engine {
	r := gin.Default()

	// v0 路由: 学生帐号服务
	v0 := r.Group("api/v0")
	{
		// 测试连接
		v0.GET("ping", api.PingPong)
	}

	// v1 路由: 管理员帐号服务
	v1 := r.Group("api/v1")
	{
		// 测试连接
		v1.GET("ping", api.PingPong)
		jw := v1.Group("jw")
		{
			// 测试教务服务连接性
			jw.GET("status", api.JWGetStatus)
			// 获取学生基础个人信息
			jw.POST("profile", api.JWGetBaseInfo)
			// 获取学生成绩
			jw.POST("score", api.JWGetScore)
		}

	}

	// v2 路由: 喜鹊儿接口 (unstable)
	v2 := r.Group("api/v3")
	{
		// 测试连接
		v2.GET("ping", api.PingPong)
	}

	return r
}
