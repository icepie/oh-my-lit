package zhyd

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/icepie/oh-my-lit/client/util"
)

// GetDormElectricity 获取寝室用电情况
func (u *ZhydUser) GetDormElectricity() (info DormElectricity, err error) {

	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://zhyd.sec.lit.edu.cn/zhyd/sydl/index", nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, cooike := range u.RealCookies {
		req.AddCookie(cooike)
	}

	log.Println(req.Cookies())

	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("User-Agent", UA)
	//req.Header.Set("Cookie", "JSESSIONID=41ED9939D8B0ED42A85A0C43AAB0D915; muyun_sign_cookie=d5f6728e5a9345832b7a4bc900dcc34a")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
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
		if info.Electricity, err = strconv.ParseFloat(electricity, 64); err == nil {
			return
		}
	}

	if len(balance) > 0 {
		if info.Balance, err = strconv.ParseFloat(balance, 64); err == nil {
			return
		}
	}

	electricitySubsidy, _ := util.GetSubstingBetweenStrings(body, `剩余补助<span class="mui-badge mui-badge-success">`, `</span></li>`)
	balanceSubsidy, _ := util.GetSubstingBetweenStrings(body, `剩余补助金额<span class="mui-badge mui-badge-success">`, `</span></li>`)

	if len(electricitySubsidy) > 0 {
		if info.ElectricitySubsidy, err = strconv.ParseFloat(electricitySubsidy, 64); err == nil {
			return
		}
	}

	if len(balanceSubsidy) > 0 {
		if info.BalanceSubsidy, err = strconv.ParseFloat(balanceSubsidy, 64); err == nil {
			return
		}
	}

	return

}
