package zhyd

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// GetDormElectricity 获取寝室用电情况
func (u *ZhydUser) GetDormElectricity() (rte []DormElectricity, err error) {

	client := &http.Client{}

	req, err := http.NewRequest("GET", GetDormElectricityURl, nil)
	if err != nil {
		return
	}

	for _, cooike := range u.RealCookies {
		req.AddCookie(cooike)
	}

	// log.Println(req.Cookies())

	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("User-Agent", UA)

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
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
				electricity := s.Text()
				if len(electricity) > 0 {
					de.Electricity, _ = strconv.ParseFloat(electricity, 64)
				}
			case 3:
				balance := s.Text()
				if len(balance) > 0 {
					de.Balance, _ = strconv.ParseFloat(balance, 64)
				}
			case 4:
				electricitySubsidy := s.Text()
				if len(electricitySubsidy) > 0 {
					de.ElectricitySubsidy, _ = strconv.ParseFloat(electricitySubsidy, 64)
				}
			case 5:
				balanceSubsidy := s.Text()
				if len(balanceSubsidy) > 0 {
					de.BalanceSubsidy, _ = strconv.ParseFloat(balanceSubsidy, 64)
				}
			}

		})
		rte = append(rte, de)
	})

	return

}

// GetElectricityDetails 获取寝室用电明细
func (u *ZhydUser) GetElectricityDetails() (rte []ElectricityDetails, err error) {

	client := &http.Client{}

	req, err := http.NewRequest("GET", GetElectricityDetailsUrl, nil)
	if err != nil {
		return
	}

	for _, cooike := range u.RealCookies {
		req.AddCookie(cooike)
	}

	// log.Println(req.Cookies())
	// }

	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("User-Agent", UA)

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	// bodyText, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return
	// }

	// body := string(bodyText)

	// log.Println(body)

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return
	}

	// log.Println(doc.Html())

	// doc.Find("div.mui-card>ul.mui-table-view>").Each(func(i int, s *goquery.Selection) {
	// 	log.Println(s.Html())
	// })

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
				electricity := s.Text()
				if len(electricity) > 0 {
					ed.Electricity, _ = strconv.ParseFloat(electricity, 64)
				}
			}
		})

		// 用电详情获取
		s.Find("div.mui-scroll>li.mui-table-view-cell").Each(func(i int, s *goquery.Selection) {

			// 去除空格
			timeStr := strings.TrimSpace(s.Nodes[0].FirstChild.Data)

			// timeStr = strings.Trim(timeStr, "\n")

			// log.Println(timeStr)

			var d Detail

			d.Time, err = time.ParseInLocation(TimeLayout, timeStr, Location)
			if err != nil {
				return
			}

			v := s.Find("span.mui-badge.mui-badge-primary").First().Text()
			if len(v) > 0 {
				d.Value, _ = strconv.ParseFloat(v, 64)
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

	client := &http.Client{}

	req, err := http.NewRequest("GET", GetChargeRecordsUrl, nil)
	if err != nil {
		return
	}

	for _, cooike := range u.RealCookies {
		req.AddCookie(cooike)
	}

	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("User-Agent", UA)

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	body := string(bodyText)

	reg := regexp.MustCompile(`this.infoList = \[(.*)\]`)

	result := reg.FindAllStringSubmatch(body, -1)

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
