package server

import (
	"github.com/gin-gonic/gin"
	"github.com/icepie/lit-edu-go/api"
	"github.com/icepie/lit-edu-go/middleware"
)

// ApiRouter 创建新路由
func ApiRouter() *gin.Engine {
	r := gin.Default()

	a := r.Group("api")
	{
		// 统一认证接口
		a.POST("auth", api.JWUserAuth)

		// v0 路由: 学生帐号服务
		v1 := a.Group("v1")
		{
			// 测试连接
			v1.GET("ping", api.PingPong)

			user := v1.Group("user")
			user.Use(middleware.JWT())
			{
				// 获取小爱课程班格式的学生个人课表
				user.GET("aischedule", api.JWAISchedule)
			}
		}

		// v1 路由: 管理员帐号服务
		v2 := a.Group("v2")
		{
			// 测试连接
			v2.GET("ping", api.PingPong)
			jw := v2.Group("jw")
			{
				// 测试教务服务连接性
				jw.GET("status", api.JWGetStatus)

			}

			user := v2.Group("user")
			user.Use(middleware.JWT())
			{
				// 获取学生基础个人信息
				user.GET("profile", api.JWGetBaseInfo)
				// 获取学生成绩
				user.GET("score", api.JWGetScore)
			}

		}

		// v3 路由: 喜鹊儿接口 (unstable)
		v3 := a.Group("v3")
		{
			// 测试连接
			v3.GET("ping", api.PingPong)
		}
	}
	return r
}
