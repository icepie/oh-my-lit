package zhyd

import (
	"time"
)

const (
	TimeLayout    = "2006/1/2 15:04:05"
	TwoTimeLayout = `"2006\/1\/2 15:04:05"`
)

var Location = time.FixedZone("GMT", 8*3600)

// DormElectricity 宿舍用电余额结构
type DormElectricity struct {
	Name               string `json:"name"`
	BuildName          string `json:"build_name"`
	Room               string `json:"room"`
	Electricity        string `json:"electricity"`
	Balance            string `json:"balance"`
	ElectricitySubsidy string `json:"electricity_subsidy"`
	BalanceSubsidy     string `json:"balance_subsidy"`
}

// ElectricityDetails 宿舍用电详情结构
type ElectricityDetails struct {
	Name        string   `json:"name"`
	BuildName   string   `json:"build_name"`
	Room        string   `json:"room"`
	Electricity string   `json:"electricity"`
	Details     []Detail `json:"details"`
}

// Detail 日详情结构
type Detail struct {
	Time  time.Time `json:"time"`
	Value string    `json:"value"`
}

// ChargeRecords 充值记录结构
type ChargeRecords struct {
	Xtbz      string `json:"XTBZ"`
	BuildName string `json:"buildName"`
	Device    string `json:"device"`
	Mdid      string `json:"mdid"`
	Mx        []struct {
		Accounttime CustomTime `json:"accounttime"`
		Inmoney     string     `json:"inmoney"`
		Paytype     string     `json:"paytype"`
	} `json:"mx"`
	Room        string `json:"room"`
	RoomID      string `json:"roomId"`
	Electricity string `json:"syl"`
}
type CustomTime struct {
	Time time.Time
}

func (ct *CustomTime) UnmarshalJSON(b []byte) error {

	body := string(b)

	// Ignore null, like in the main JSON package.
	if body == "null" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	var err error
	ct.Time, err = time.ParseInLocation(TwoTimeLayout, body, Location)
	return err
}
