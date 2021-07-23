package health

// LoginParam 登陆参数
type LoginParam struct {
	CardNo   string `json:"cardNo"`
	Password string `json:"password"`
}

// Result 响应结果
type Result struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// UserInfo 用户信息
type UserInfo struct {
	Age                    int         `json:"age"`
	CardNo                 interface{} `json:"cardNo"`
	ExpireTime             string      `json:"expireTime"`
	Identity               int         `json:"identity"`
	IsAdmin                int         `json:"isAdmin"`
	IsApprover             interface{} `json:"isApprover"`
	IsGeneralAdmin         int         `json:"isGeneralAdmin"`
	IsReportAdmin          int         `json:"isReportAdmin"`
	IsReturnSchoolApprover int         `json:"isReturnSchoolApprover"`
	IsTwoTemperature       int         `json:"isTwoTemperature"`
	LastUpdateTime         string      `json:"lastUpdateTime"`
	LocalAddress           string      `json:"localAddress"`
	LogoURL                string      `json:"logoUrl"`
	Mobile                 string      `json:"mobile"`
	Name                   string      `json:"name"`
	NativePlaceAddress     string      `json:"nativePlaceAddress"`
	NativePlaceCity        string      `json:"nativePlaceCity"`
	NativePlaceDistrict    string      `json:"nativePlaceDistrict"`
	NativePlaceProvince    string      `json:"nativePlaceProvince"`
	OrganizationName       string      `json:"organizationName"`
	Sex                    int         `json:"sex"`
	TeamCity               string      `json:"teamCity"`
	TeamID                 int         `json:"teamId"`
	TeamName               string      `json:"teamName"`
	TeamNo                 string      `json:"teamNo"`
	TeamProvince           string      `json:"teamProvince"`
	Token                  string      `json:"token"`
	UserID                 int         `json:"userId"`
	UserOrganizationID     int         `json:"userOrganizationId"`
}

type LastRecord struct {
	AbroadInfo               string      `json:"abroadInfo"`
	CaseAddress              interface{} `json:"caseAddress"`
	ContactAddress           string      `json:"contactAddress"`
	ContactCity              string      `json:"contactCity"`
	ContactDistrict          string      `json:"contactDistrict"`
	ContactPatient           string      `json:"contactPatient"`
	ContactProvince          string      `json:"contactProvince"`
	ContactTime              interface{} `json:"contactTime"`
	CreateTime               string      `json:"createTime"`
	CureTime                 interface{} `json:"cureTime"`
	CurrentAddress           string      `json:"currentAddress"`
	CurrentCity              string      `json:"currentCity"`
	CurrentDistrict          string      `json:"currentDistrict"`
	CurrentProvince          string      `json:"currentProvince"`
	CurrentStatus            string      `json:"currentStatus"`
	DiagnosisTime            interface{} `json:"diagnosisTime"`
	ExceptionalCase          int64       `json:"exceptionalCase"`
	ExceptionalCaseInfo      string      `json:"exceptionalCaseInfo"`
	FriendHealthy            int64       `json:"friendHealthy"`
	GoHuBeiCity              string      `json:"goHuBeiCity"`
	GoHuBeiTime              interface{} `json:"goHuBeiTime"`
	HealthyStatus            int64       `json:"healthyStatus"`
	ID                       int64       `json:"id"`
	IsAbroad                 int64       `json:"isAbroad"`
	IsInTeamCity             int64       `json:"isInTeamCity"`
	Isolation                int64       `json:"isolation"`
	PeerAddress              interface{} `json:"peerAddress"`
	PeerIsCase               int64       `json:"peerIsCase"`
	ReportDate               string      `json:"reportDate"`
	SeekMedical              int64       `json:"seekMedical"`
	SeekMedicalInfo          string      `json:"seekMedicalInfo"`
	SelfHealthy              int64       `json:"selfHealthy"`
	SelfHealthyInfo          string      `json:"selfHealthyInfo"`
	SelfHealthyTime          interface{} `json:"selfHealthyTime"`
	TeamID                   int64       `json:"teamId"`
	Temperature              string      `json:"temperature"`
	TemperatureNormal        int64       `json:"temperatureNormal"`
	TemperatureThree         string      `json:"temperatureThree"`
	TemperatureTwo           string      `json:"temperatureTwo"`
	TravelPatient            string      `json:"travelPatient"`
	TreatmentHospitalAddress string      `json:"treatmentHospitalAddress"`
	UserID                   int64       `json:"userId"`
	VillageIsCase            int64       `json:"villageIsCase"`
}
