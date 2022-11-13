package zhyd

import (
	"bytes"
	"encoding/json"
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// GetDormElectricity 获取寝室用电情况
func (u *ZhydUser) GetDormElectricity() (rte []DormElectricity, err error) {

	resp, _ := u.Client.R().
		Get(GetDormElectricityURl)

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(resp.Body()))
	if err != nil {
		return
	}

	// 获取基本信息
	doc.Find("div.mui-card>ul.mui-table-view").Each(func(i int, s *goquery.Selection) {

		var de DormElectricity

		name := s.Find("li.mui-table-view-divider").First()
		de.Name = name.Text()

		s.Find("span.mui-badge").Each(func(i int, s *goquery.Selection) {
			switch i {
			case 0:
				de.BuildName = s.Text()
			case 1:
				de.Room = s.Text()
			case 2:
				de.Electricity = s.Text()
			case 3:
				de.Balance = s.Text()
			case 4:
				de.ElectricitySubsidy = s.Text()
			case 5:
				de.BalanceSubsidy = s.Text()
			}

		})
		rte = append(rte, de)
	})

	return

}

// GetElectricityDetails 获取寝室用电明细
func (u *ZhydUser) GetElectricityDetails() (rte []ElectricityDetails, err error) {

	resp, _ := u.Client.R().
		Get(GetElectricityDetailsUrl)

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(resp.Body()))
	if err != nil {
		return
	}

	doc.Find("div.mui-content>div").Each(func(i int, s *goquery.Selection) {

		var ed ElectricityDetails

		name := s.Find("div.mui-card>ul.mui-table-view>li.mui-table-view-divider").First()
		ed.Name = name.Text()

		s.Find("div.mui-card>ul.mui-table-view>li>span.mui-badge").Each(func(i int, s *goquery.Selection) {
			switch i {
			case 0:
				ed.BuildName = s.Text()
			case 1:
				ed.Room = s.Text()
			case 2:
				ed.Electricity = s.Text()
			}
		})

		// 用电详情获取
		s.Find("div.mui-scroll>li.mui-table-view-cell").Each(func(i int, s *goquery.Selection) {

			// 去除空格
			timeStr := strings.TrimSpace(s.Nodes[0].FirstChild.Data)

			// timeStr = strings.Trim(timeStr, "\n")

			// log.Println(timeStr)

			var d Detail

			dTime, err := time.ParseInLocation(TimeLayout, timeStr, Location)
			if err != nil {
				return
			}

			d.Time = dTime.Format("2006-01-02 15:04:05")

			v := s.Find("span.mui-badge.mui-badge-primary").First().Text()
			if len(v) > 0 {
				d.Value = v
			}

			ed.Details = append(ed.Details, d)

		})

		// 添加到最后结果
		rte = append(rte, ed)
	})

	return

}

// GetChargeRecords 获取消费记录
func (u *ZhydUser) GetChargeRecords() (rte []ChargeRecords, err error) {

	resp, _ := u.Client.R().
		Get(GetChargeRecordsUrl)

	reg := regexp.MustCompile(`this.infoList = \[(.*)\]`)

	result := reg.FindAllStringSubmatch(resp.String(), -1)

	if len(result) == 0 {
		err = errors.New("no result")
		return
	}

	err = json.Unmarshal([]byte("["+result[0][1]+"]"), &rte)
	if err != nil {
		return
	}

	return

}
