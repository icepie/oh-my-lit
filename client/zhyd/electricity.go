package zhyd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// GetDormElectricity 获取寝室用电情况
func (u *ZhydUser) GetDormElectricity() {

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
	fmt.Printf("%s\n", bodyText)

}
