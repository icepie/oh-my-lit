package service

import (
	"encoding/json"
	"strconv"

	"github.com/asmcos/requests"
	"github.com/icepie/lit-edu-go/model"
	"github.com/icepie/lit-edu-go/pkg/e"
	"github.com/icepie/lit-edu-go/service/jw"
	log "github.com/sirupsen/logrus"
)

const (
	AIScheduleURL  = "https://open-schedule.ai.xiaomi.com/api/schedule/parser"
	AIScheduleTBID = "38307"
)

type AIScheduleReq struct {
	Html string `json:"html"`
	TBID string `json:"tb_id"`
}

type AIScheduleRes struct {
	TBID      string `json:"tb_id"`
	ParserRet string `json:"parserRet"`
}

type AISchedule struct {
	CourseInfos []struct {
		Day      int    `json:"day"`
		Name     string `json:"name"`
		Position string `json:"position"`
		Sections []struct {
			Section int `json:"section"`
		} `json:"sections"`
		Teacher string `json:"teacher"`
		Weeks   []int  `json:"weeks"`
	} `json:"courseInfos"`
	SectionTimes []struct {
		EndTime   string `json:"endTime"`
		Section   int    `json:"section"`
		StartTime string `json:"startTime"`
	} `json:"sectionTimes"`
}

// GetAIscheduleService 获取小爱课程表格式服务结构
type GetAIscheduleService struct {
	User model.StuAccount
}

// GetBaseInfo 根据 StuID 获取学生基本信息
func (service *GetAIscheduleService) GetAIschedule() model.Response {

	if service.User.PassWord == "" {
		code := e.Error
		return model.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "please enter the correct password",
		}
	}

	cookies, err := jw.SendLogin(service.User.StuID, service.User.PassWord, "STU")
	if err != nil {
		log.Warningln(err)
		code := e.Error
		return model.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	body, err := jw.QuerySchedule(cookies)
	if err != nil {
		log.Warningln(err)
		code := e.Error
		return model.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	req := AIScheduleReq{Html: strconv.Quote(string(body)), TBID: AIScheduleTBID}

	reqjson, err := json.Marshal(req)
	if err != nil {
		log.Warningln(err)
		code := e.Error
		return model.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	resp, err := requests.PostJson(AIScheduleURL, string(reqjson))
	if err != nil {
		log.Warningln(err)
		code := e.Error
		return model.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	respjson := AIScheduleRes{}
	err = json.Unmarshal([]byte(resp.Text()), &respjson)
	if err != nil {
		log.Warningln(err)
		code := e.Error
		return model.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	schedule := AISchedule{}

	err = json.Unmarshal([]byte(respjson.ParserRet), &schedule)
	if err != nil {
		log.Warningln(err)
		code := e.Error
		return model.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	code := e.Success
	return model.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   schedule,
	}
}
