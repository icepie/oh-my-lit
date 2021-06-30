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
	BuildName          string
	Room               string
	Electricity        float64
	Balance            float64
	ElectricitySubsidy float64
	BalanceSubsidy     float64
}

// ElectricityDetails 宿舍用电详情结构
type ElectricityDetails struct {
	Building    string
	Room        string
	Electricity float64
	Details     []Detail
}

// Detail 日详情结构
type Detail struct {
	Time  time.Time
	Value float64
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
