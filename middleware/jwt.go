package middleware

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/icepie/lit-edu-go/model"
	"github.com/icepie/lit-edu-go/pkg/e"
	"github.com/icepie/lit-edu-go/pkg/util"

	"github.com/gin-gonic/gin"
)

// JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = 200
		token := c.Query("token")
		if token == "" {
			code = 400
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ErrorJWTCheck
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ErrorJWTTimeout
			}

			log.Println(claims)

			// token data 解密
			info := util.UserDecrypt(claims.Info)

			userRaw := strings.Split(info, ",")

			if len(userRaw) != 2 {
				code = e.ErrorJWTCheck
			} else {
				user := model.StuAccount{StuID: userRaw[0], PassWord: userRaw[1]}
				c.Set("user", user)
			}

		}

		if code != e.Success {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
