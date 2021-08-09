package health

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

// FirstReportByRaw 使用原始参数进行今日第一次上报
func (u *HealthUser) FirstReportByRaw(param FirstRecordParam) (err error) {

	body, err := json.Marshal(param)
	if err != nil {
		return
	}

	resp, err := u.Client.R().
		SetBody(string(body)).
		SetResult(Result{}).
		Post(FirstRecordUrl)

	if err != nil {
		return
	}

	r := resp.Result().(*Result)

	// 登陆失败
	if r.Code != 200 {
		err = errors.New(r.Msg)
	}

	return
}

// FirstReport 进行今日第一次上报
func (u *HealthUser) FirstReport(firstTemp float64, secondTemp float64, thirdTemp float64) (err error) {

	lr, err := u.GetLastRecord()
	if err != nil {
		return
	}

	// 构建上报时间 (北京时间东八区)
	var cstZone = time.FixedZone("CST", 8*3600) // 东八区
	nowTime := time.Now().In(cstZone)

	param := FirstRecordParam{
		AbroadInfo:               lr.AbroadInfo,
		CaseAddress:              lr.CaseAddress,
		ContactAddress:           lr.ContactAddress,
		ContactCity:              lr.ContactCity,
		ContactDistrict:          lr.ContactDistrict,
		ContactLocation:          "",
		ContactPatient:           lr.ContactPatient,
		ContactProvince:          lr.ContactProvince,
		ContactTime:              lr.ContactTime,
		CureTime:                 lr.CureTime,
		CurrentAddress:           lr.CurrentAddress,
		CurrentCity:              lr.ContactCity,
		CurrentDistrict:          lr.CurrentDistrict,
		CurrentLocation:          "",
		CurrentProvince:          lr.CurrentProvince,
		CurrentStatus:            lr.CurrentStatus,
		DiagnosisTime:            lr.DiagnosisTime,
		ExceptionalCase:          lr.ExceptionalCase,
		ExceptionalCaseInfo:      lr.ExceptionalCaseInfo,
		FriendHealthy:            lr.FriendHealthy,
		GoHuBeiCity:              lr.GoHuBeiCity,
		GoHuBeiTime:              lr.GoHuBeiTime,
		HealthyStatus:            lr.HealthyStatus,
		IsAbroad:                 lr.IsAbroad,
		IsInTeamCity:             lr.IsInTeamCity,
		IsTrip:                   lr.IsAbroad,
		Isolation:                lr.Isolation,
		LocalAddress:             u.UserInfo.LocalAddress,
		Mobile:                   u.UserInfo.Mobile,
		NativePlaceAddress:       u.UserInfo.NativePlaceAddress,
		NativePlaceCity:          u.UserInfo.NativePlaceCity,
		NativePlaceDistrict:      u.UserInfo.NativePlaceDistrict,
		NativePlaceProvince:      u.UserInfo.NativePlaceProvince,
		PeerAddress:              lr.PeerAddress,
		PeerIsCase:               lr.PeerIsCase,
		ReportDate:               nowTime.Format("2006-01-02"),
		SeekMedical:              lr.SeekMedical,
		SeekMedicalInfo:          lr.SeekMedicalInfo,
		SelfHealthy:              lr.SelfHealthy,
		SelfHealthyInfo:          lr.SelfHealthyInfo,
		SelfHealthyTime:          lr.SelfHealthyTime,
		TeamID:                   lr.TeamID,
		Temperature:              fmt.Sprint(secondTemp),
		TemperatureNormal:        lr.TemperatureNormal,
		TemperatureThree:         fmt.Sprint(secondTemp),
		TemperatureTwo:           fmt.Sprint(thirdTemp),
		TravelPatient:            lr.TravelPatient,
		TreatmentHospitalAddress: lr.TreatmentHospitalAddress,
		UserID:                   u.UserInfo.UserID,
		VillageIsCase:            lr.VillageIsCase,
	}

	err = u.FirstReportByRaw(param)

	return
}

// SecondReportByRaw 使用原始参数进行今日第二次上报
func (u *HealthUser) SecondReportByRaw(param ExtraRecordParam) (err error) {

	body, err := json.Marshal(param)
	if err != nil {
		return
	}

	resp, err := u.Client.R().
		SetBody(string(body)).
		SetResult(Result{}).
		Put(SecondRecordUrl)

	if err != nil {
		return
	}

	r := resp.Result().(*Result)

	// 登陆失败
	if r.Code != 200 {
		err = errors.New(r.Msg)
	}

	return
}

// SecondReport 进行今日第二次上报
func (u *HealthUser) SecondReport(temp float64) (err error) {

	lr, err := u.GetLastRecord()
	if err != nil {
		return
	}

	param := ExtraRecordParam{
		HealthyRecordID:   lr.ID,
		Temperature:       fmt.Sprint(temp),
		TemperatureNormal: lr.TemperatureNormal,
	}

	err = u.SecondReportByRaw(param)

	return

}

// ThirdReportByRaw 使用原始参数进行今日第三次上报
func (u *HealthUser) ThirdReportByRaw(param ExtraRecordParam) (err error) {

	body, err := json.Marshal(param)
	if err != nil {
		return
	}

	resp, err := u.Client.R().
		SetBody(string(body)).
		SetResult(Result{}).
		Put(ThirdRecordUrl)

	if err != nil {
		return
	}

	r := resp.Result().(*Result)

	// 登陆失败
	if r.Code != 200 {
		err = errors.New(r.Msg)
	}

	return
}

// ThirdReport 进行今日第三次上报
func (u *HealthUser) ThirdReport(temp float64) (err error) {

	lr, err := u.GetLastRecord()
	if err != nil {
		return
	}

	param := ExtraRecordParam{
		HealthyRecordID:   lr.ID,
		Temperature:       fmt.Sprint(temp),
		TemperatureNormal: lr.TemperatureNormal,
	}

	err = u.ThirdReportByRaw(param)

	return

}
