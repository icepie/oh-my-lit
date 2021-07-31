package zhyd

import (
	"github.com/go-resty/resty/v2"
)

var (
	// AuthHostUrl 登陆主页面
	AuthHostUrl = "http://ids.lit.edu.cn"
	// NeedCaptchaUrl 判断是否需要验证码的接口
	NeedCaptchaUrl = AuthHostUrl + "/authserver/needCaptcha.html"
	// CaptchaUrl 获取验证码
	CaptchaUrl = AuthHostUrl + "/authserver/captcha.html"
	// LoginUrl 登陆接口
	LoginUrl = AuthHostUrl + "/authserver/login"
	// ZhydHost
	ZhydHost = "http://zhyd.sec.lit.edu.cn"
	// ZhydHostUrl 智慧用电主页
	ZhydHostUrl = ZhydHost + "/zhyd"
	// GetDormElectricityURl 获取宿舍电量主页
	GetDormElectricityURl = ZhydHostUrl + "/sydl/index"
	// GetElectricityDetailsUrl 获取用电明细
	GetElectricityDetailsUrl = ZhydHostUrl + "/ydmx/index"
	// GetConsumptionRecordsUrl 获取消费记录
	GetChargeRecordsUrl = ZhydHostUrl + "/zzgd/index"
	// UA
	UA = "User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36"
	// MainHeaders 主请求头
	MainHeaders = map[string]string{
		"User-Agent":      UA,
		"Accept":          "*/*",
		"Accept-Encoding": "gzip, deflate",
		"Connection":      "keep-alive",
	}
)

// ZhydUser 智能控电用户结构体
type ZhydUser struct {
	Username    string
	Password    string
	Client      *resty.Client
}

// SetPassword 设置用户名
func (u *ZhydUser) SetUsername(username string) *ZhydUser {
	u.Username = username
	return u
}

// SetPassword 设置密码
func (u *ZhydUser) SetPassword(password string) *ZhydUser {
	u.Password = password
	return u
}

// NewZhydUser 新建智能控电用户
func NewZhydUser() *ZhydUser {

	var u ZhydUser

	u.Client = resty.New()
	u.Client.SetHeaders(MainHeaders)

	// 拿个cookies
	u.PerSetCooikes()

	return &u
}

// PerSetCooikes 预先设置必要Cookies
func (u *ZhydUser) PerSetCooikes() *ZhydUser {

	// 先访问一下页面，获取cookie
	u.Client.R().
		Get(ZhydHost)

	return u
}
