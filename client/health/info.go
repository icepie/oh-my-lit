package health

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"
)

// GetLastRecord 获取上次上报记录
func (u *HealthUser) GetLastRecord() (lastRecord LastRecord, err error) {

	resp, err := u.Client.R().
		SetQueryParams(map[string]string{
			"teamId": strconv.Itoa(u.UserInfo.TeamID),
			"userId": strconv.Itoa(u.UserInfo.UserID),
		}).SetResult(Result{}).
		Get(LastRecordUrl)

	if err != nil {
		return
	}

	r := resp.Result().(*Result)

	// 登陆失败
	if r.Code != 200 {
		err = errors.New(r.Msg)
		return
	}

	byteData, _ := json.Marshal(r.Data)
	err = json.Unmarshal(byteData, &lastRecord)
	if err != nil {
		return
	}

	return
}

// IsReportedToday 今日是否上报
// rtime: 次数 [0,1,2,3] 0: 为三次是否全上报 1: 第一次 2: 第二次  3: 第三次
func (u *HealthUser) IsReportedToday(rtime uint) (isReported bool, err error) {

	isReported = false

	if rtime > 3 {
		err = errors.New("rtime error")
		return
	}

	lastRecord, err := u.GetLastRecord()
	if err != nil {
		return
	}

	// 查询最近上报时间是否为今天 (北京时间东八区)
	var cstZone = time.FixedZone("CST", 8*3600) // 东八区
	nowTime := time.Now().In(cstZone)

	createTime, err := time.Parse(TimeLayout, lastRecord.CreateTime)
	if err != nil {
		return
	}

	if createTime.Format(TimeLayout2) == nowTime.Format(TimeLayout2) {

		if rtime == 0 {
			if len(lastRecord.Temperature) != 0 && len(lastRecord.TemperatureTwo) != 0 && len(lastRecord.TemperatureThree) != 0 {
				isReported = true
			}

		} else if rtime == 1 {
			if len(lastRecord.Temperature) != 0 {
				isReported = true
			}
		} else if rtime == 2 {
			if len(lastRecord.TemperatureTwo) != 0 {
				isReported = true
			}
		} else if len(lastRecord.TemperatureThree) != 0 {
			isReported = true
		}

		return
	}

	return
}
