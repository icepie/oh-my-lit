package service

import (
	"github.com/icepie/lit-edu-go/model"
	"github.com/icepie/lit-edu-go/pkg/e"
	"github.com/icepie/lit-edu-go/service/jw"
)

// GetStatusService 获取教务管理员帐号登陆情况以及其他信息
type GetStatusService struct {
}

// GetStatus 根据 StuID 获取学生基本信息
func (service *GetStatusService) GetStatus() model.Response {

	if iswork, err := jw.IsWork(); !iswork || err != nil {
		return model.Response{
			Status: 200,
			Data:   "",
			Msg:    "please use the right sys account of jw",
			Error:  err.Error(),
		}
	}

	code := e.SUCCESS
	return model.Response{
		Status: code,
		Msg:    "lit jw is work fine!",
	}
}
