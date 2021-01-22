package service

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/djimenez/iconv-go"
)

func SendLogin() {
	res, err := http.Get("http://jw.sec.lit.edu.cn/_data/index_LOGIN.aspx")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	utf8Body, err := iconv.NewReader(res.Body, "gb2312", "utf-8")
	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(utf8Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("input[name=__VIEWSTATE]").Each(func(i int, s *goquery.Selection) {
		vs, ex := s.Attr("value")
		if ex {
			println(vs)
		}
	})
}
