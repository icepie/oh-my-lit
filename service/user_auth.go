package service

import (
	"github.com/icepie/lit-edu-go/model"
	"github.com/icepie/lit-edu-go/pkg/e"
	"github.com/icepie/lit-edu-go/pkg/util"
	"github.com/icepie/lit-edu-go/service/jw"
	log "github.com/sirupsen/logrus"
)

// UserAuthService 用户认证
type UserAuthService struct {
	StuID    string `json:"stuid" binding:"required"`
	PassWord string `json:"password" binding:"required"`
}

// UserAuth 用户认证并获取token
func (service *UserAuthService) UserAuth() model.Response {

	if service.PassWord == "" {
		code := e.Error
		return model.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "please enter the correct password",
		}
	}

	_, err := jw.SendLogin(service.StuID, service.PassWord, "STU")
	if err != nil {
		log.Warningln(err)
		code := e.Error
		return model.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 更新 Token
	token, err := util.GenerateToken(util.UserEncrypt(service.StuID, service.PassWord))
	if err != nil {
		log.Warningln(err)
		code := e.ErrorGenToken
		return model.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	log.Info("token: ", token)

	code := e.Success
	return model.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   model.TokenData{Token: token},
	}

}
