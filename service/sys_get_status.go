package service

import (
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/thinkeridea/go-extend/exunicode/exutf8"

	"github.com/PuerkitoBio/goquery"
	"github.com/icepie/lit-edu-go/model"
	"github.com/icepie/lit-edu-go/pkg/e"
	"github.com/icepie/lit-edu-go/service/jw"
)

// GetStatusService 获取教务管理员帐号登陆情况以及其他信息
type GetStatusService struct {
}

// GetStatus 根据 StuID 获取学生基本信息
func (service *GetStatusService) GetStatus() model.Response {

	var status model.JWStatus

	if iswork, err := jw.IsWork(); !iswork || err != nil {
		return model.Response{
			Status: 200,
			Data:   "",
			Msg:    "please use the right sys account of jw",
			Error:  err.Error(),
		}
	}

	body, err := jw.GetBanner(jw.JWCookies)
	if err != nil {
		log.Warningln(err)
		code := e.Error
		return model.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		log.Warningln(err)
		code := e.Error
		return model.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 在线人数
	onstrfull := doc.Find("span").First().Text()

	// 去掉缩进和空格
	onstrq := strings.Replace(strings.Replace(strings.Replace(onstrfull, "\\t", "", -1), "\n", "", -1), "\t", "", -1)
	// 去掉特殊的符号
	onstrq = strings.Replace(onstrq, "\u00a0", "", -1)
	// 截取一下
	onstr := exutf8.RuneSubString(onstrq, 5, 10)
	// 尝试转为整形
	on, err := strconv.Atoi(onstr)

	if err == nil {
		status.OnlineNumber = on
	} else {
		status.OnlineNumber = onstr
	}

	code := e.Success
	return model.Response{
		Status: code,
		Msg:    "lit jw is work fine!",
		Data:   status,
	}
}
