package zhyd

import "time"

// DormElectricity 宿舍用电余额结构
type DormElectricity struct {
	Building           string
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
