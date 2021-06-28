package zhyd

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"

	"github.com/icepie/oh-my-lit/client/util"
)

const (
	TimeLayout = "2006/1/2 15:04:05"
)

// GetDormElectricity 获取寝室用电情况
func (u *ZhydUser) GetDormElectricity() (info DormElectricity, err error) {

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
	//req.Header.Set("Cookie", "JSESSIONID=41ED9939D8B0ED42A85A0C43AAB0D915; muyun_sign_cookie=d5f6728e5a9345832b7a4bc900dcc34a")
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

	if !strings.Contains(body, "剩余用电") {
		err = errors.New("fail to get info")
		return
	}

	info.Building, _ = util.GetSubstingBetweenStrings(body, `绑定楼栋<span class="mui-badge mui-badge-primary">`, `</span></li>`)
	info.Room, _ = util.GetSubstingBetweenStrings(body, `绑定房间<span class="mui-badge mui-badge-primary">`, `</span></li>`)
	electricity, _ := util.GetSubstingBetweenStrings(body, `剩余电量<span class="mui-badge mui-badge-success">`, `</span></li>`)
	balance, _ := util.GetSubstingBetweenStrings(body, `剩余金额<span class="mui-badge mui-badge-success">`, `</span></li>`)

	if len(electricity) > 0 {
		info.Electricity, _ = strconv.ParseFloat(electricity, 64)
	}

	if len(balance) > 0 {
		info.Balance, err = strconv.ParseFloat(balance, 64)
	}

	electricitySubsidy, _ := util.GetSubstingBetweenStrings(body, `剩余补助<span class="mui-badge mui-badge-success">`, `</span></li>`)
	balanceSubsidy, _ := util.GetSubstingBetweenStrings(body, `剩余补助金额<span class="mui-badge mui-badge-success">`, `</span></li>`)

	if len(electricitySubsidy) > 0 {
		info.ElectricitySubsidy, _ = strconv.ParseFloat(electricitySubsidy, 64)
	}

	if len(balanceSubsidy) > 0 {
		info.BalanceSubsidy, _ = strconv.ParseFloat(balanceSubsidy, 64)
	}

	return

}

// GetElectricityDetails 获取寝室用电明细
func (u *ZhydUser) GetElectricityDetails() (rte ElectricityDetails, err error) {

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
	//req.Header.Set("Cookie", "JSESSIONID=41ED9939D8B0ED42A85A0C43AAB0D915; muyun_sign_cookie=d5f6728e5a9345832b7a4bc900dcc34a")
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

	// rte.Building = doc.Find("span.mui-badge.mui-badge-primary").First().Text()

	// rte.Room = doc.Find("span.mui-badge.mui-badge-primary").Eq(1).Text()

	// electricity := doc.Find("span.mui-badge.mui-badge-primary").Eq(2).Text()

	// 获取基本信息
	doc.Find("span.mui-badge.mui-badge-primary").Each(func(i int, s *goquery.Selection) {
		switch i {
		case 0:
			rte.Building = s.Text()
		case 1:
			rte.Room = s.Text()
		case 2:
			electricity := s.Text()
			if len(electricity) > 0 {
				rte.Electricity, _ = strconv.ParseFloat(electricity, 64)
			}
			// default:

		}
	})

	// 用电详情获取
	doc.Find("div.mui-scroll>li.mui-table-view-cell").Each(func(i int, s *goquery.Selection) {

		// 去除空格
		timeStr := strings.TrimSpace(s.Nodes[0].FirstChild.Data)

		// timeStr = strings.Trim(timeStr, "\n")

		// log.Println(timeStr)

		var d Detail

		d.Time, err = time.Parse(TimeLayout, timeStr)
		if err != nil {
			return
		}

		v := s.Find("span.mui-badge.mui-badge-primary").First().Text()
		if len(v) > 0 {
			d.Value, _ = strconv.ParseFloat(v, 64)
		}

		rte.Details = append(rte.Details, d)

	})

	return

}
