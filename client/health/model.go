package health

// LoginParam 登陆参数
type LoginParam struct {
	CardNo   string `json:"cardNo"`
	Password string `json:"password"`
}

// Result 响应结果
type Result struct {
	Success bool        `json:"success"`
	Code    int64       `json:"code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// UserInfo 用户信息
type UserInfo struct {
	Age                    int64       `json:"age"`
	CardNo                 interface{} `json:"cardNo"`
	ExpireTime             string      `json:"expireTime"`
	Identity               int64       `json:"identity"`
	IsAdmin                int64       `json:"isAdmin"`
	IsApprover             interface{} `json:"isApprover"`
	IsGeneralAdmin         int64       `json:"isGeneralAdmin"`
	IsReportAdmin          int64       `json:"isReportAdmin"`
	IsReturnSchoolApprover int64       `json:"isReturnSchoolApprover"`
	IsTwoTemperature       int64       `json:"isTwoTemperature"`
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
	Sex                    int64       `json:"sex"`
	TeamCity               string      `json:"teamCity"`
	TeamID                 int64       `json:"teamId"`
	TeamName               string      `json:"teamName"`
	TeamNo                 string      `json:"teamNo"`
	TeamProvince           string      `json:"teamProvince"`
	Token                  string      `json:"token"`
	UserID                 int64       `json:"userId"`
	UserOrganizationID     int64       `json:"userOrganizationId"`
}

// LastRecord 上次上报结构
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

// FirstRecordParam 首次上报参数结构
type FirstRecordParam struct {
	AbroadInfo               string      `json:"abroadInfo"`
	CaseAddress              interface{} `json:"caseAddress"`
	ContactAddress           string      `json:"contactAddress"`
	ContactCity              string      `json:"contactCity"`
	ContactDistrict          string      `json:"contactDistrict"`
	ContactLocation          string      `json:"contactLocation"`
	ContactPatient           string      `json:"contactPatient"`
	ContactProvince          string      `json:"contactProvince"`
	ContactTime              interface{} `json:"contactTime"`
	CureTime                 interface{} `json:"cureTime"`
	CurrentAddress           string      `json:"currentAddress"`
	CurrentCity              string      `json:"currentCity"`
	CurrentDistrict          string      `json:"currentDistrict"`
	CurrentLocation          string      `json:"currentLocation"`
	CurrentProvince          string      `json:"currentProvince"`
	CurrentStatus            string      `json:"currentStatus"`
	DiagnosisTime            interface{} `json:"diagnosisTime"`
	ExceptionalCase          int64       `json:"exceptionalCase"`
	ExceptionalCaseInfo      string      `json:"exceptionalCaseInfo"`
	FriendHealthy            int64       `json:"friendHealthy"`
	GoHuBeiCity              string      `json:"goHuBeiCity"`
	GoHuBeiTime              interface{} `json:"goHuBeiTime"`
	HealthyStatus            int64       `json:"healthyStatus"`
	IsAbroad                 int64       `json:"isAbroad"`
	IsInTeamCity             int64       `json:"isInTeamCity"`
	IsTrip                   int64       `json:"isTrip"`
	Isolation                int64       `json:"isolation"`
	LocalAddress             string      `json:"localAddress"`
	Mobile                   string      `json:"mobile"`
	NativePlaceAddress       string      `json:"nativePlaceAddress"`
	NativePlaceCity          string      `json:"nativePlaceCity"`
	NativePlaceDistrict      string      `json:"nativePlaceDistrict"`
	NativePlaceProvince      string      `json:"nativePlaceProvince"`
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

// ExtraRecordParam 补充上报参数 (第二, 三次)
type ExtraRecordParam struct {
	HealthyRecordID   int64  `json:"healthyRecordId"`
	Temperature       string `json:"temperature"`
	TemperatureNormal int64  `json:"temperatureNormal"`
}

// HealthyReportCount 上报统计结构
type HealthyReportCount struct {
	NotReported int64 `json:"notReported"`
	RecordTotal int64 `json:"recordTotal"`
	Reported    int64 `json:"reported"`
}
