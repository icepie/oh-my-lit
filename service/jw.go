package service

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

func FetchHomePage() {
	// Create a Resty Client
	client := resty.New()

	resp, err := client.R().
		SetHeaders(map[string]string{
			"Host":                      "jw.sec.lit.edu.cn",
			"Connection":                "keep-alive",
			"Cache-Control":             "max-age=0",
			"Upgrade-Insecure-Requests": "1",
			"Cooikes":                   "name=value; name=value; ASP.NET_SessionId=vujnou45dn4r4e3phf0lgxyn",
			"User-Agent":                "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Safari/537.36",
			"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3",
			"Referer":                   "http://120.194.42.205:9001/_data/index_LOGIN.aspx",
			"Accept-Encoding":           "gzip, deflate",
			"Accept-Language":           "zh-CN,zh;q=0.9",
		}).
		Get("http://jw.sec.lit.edu.cn/MAINFRM.aspx")

	// Explore response object
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()
}
