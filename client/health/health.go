package health

import (
	"errors"
	"time"

	"github.com/go-resty/resty/v2"
)

var (
	// MianHost
	MianHost = "hmgr.sec.lit.edu.cn"
	// MianUrl 主网址
	MianUrl = "http://" + MianHost + "/wms"
	// MuyunUrl 木云反代
	MuyunUrl = MianUrl + "/web/"
	// LoginUrl 登录接口
	LoginUrl = MianUrl + "/healthyLogin"
	// LastRecordUrl 获取上次记录
	LastRecordUrl = MianUrl + "/lastHealthyRecord"
	// FirstRecordUrl 第一次上报
	FirstRecordUrl = MianUrl + "/addHealthyRecord"
	// SecondRecordUrl 第二次上报
	SecondRecordUrl = MianUrl + "/addTwoHealthyRecord"
	// ThirdRecordUrl 第三次上报
	ThirdRecordUrl = MianUrl + "/addThreeHealthyRecord"
	// MainHeaders 主请求头
	MainHeaders = map[string]string{
		"Proxy-Connection": "keep-alive",
		"Accept":           "application/json, text/plain, */*",
		"DNT":              "1",
		"User-Agent":       "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.96 Safari/537.36",
		"Content-Type":     "application/json;charset=UTF-8;Access-Control-Allow-Headers",
		"Origin":           MianUrl,
		"Referer":          MuyunUrl,
		"Accept-Language":  "zh-CN,zh;q=0.9",
	}

	// IdentityNameList 身份
	IdentityNameList = map[string]int{
		"student": 1000401,
		"teacher": 1000402,
		"other":   1000405,
	}

	// IdentityCodeList 身份
	IdentityList = map[int]string{
		1000401: "student",
		1000402: "teacher",
		1000405: "other",
	}

	// TimeLayout 时间格式
	TimeLayout = "2006-01-02 15:04:05"
	// TimeLayout2 时间格式
	TimeLayout2 = "2006-01-02"
)

// HealthUser 健康平台用户结构体
type HealthUser struct {
	Username string
	Password string
	UserInfo UserInfo
	Client   *resty.Client // 测试用
}

// NewHealthUser 新建健康平台用户
func NewHealthUser() *HealthUser {

	var u HealthUser

	u.Client = resty.New()
	// u.Username = Username
	// u.Password = Password

	// u.Client.SetDebug(true)
	u.Client.SetHeaders(MainHeaders)
	u.Client.SetTimeout(5 * time.Second)

	return &u
}

// SetPassword 设置用户名
func (u *HealthUser) SetUsename(username string) *HealthUser {
	u.Username = username
	return u
}

// SetPassword 设置密码
func (u *HealthUser) SetPassword(password string) *HealthUser {
	u.Password = password
	return u
}

// PerSetCooikes 访问一下反代页面
func (u *HealthUser) PerSetCooikes() (err error) {

	// 先访问一下反代页面，获取cookie
	resp, err := u.Client.R().
		Get(MuyunUrl)

	if err != nil {
		return
	}

	if resp.IsError() {
		err = errors.New("fail to new the object")
		return
	}

	return
}
