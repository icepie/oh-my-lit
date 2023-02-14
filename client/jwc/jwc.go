package jwc

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"
)

const (
	// ggtz
	GGTZ_URL = "https://www.lit.edu.cn/jwc/xb2021/ggtz.htm"
	// gttz by page
	GGTZ_PAGE_URL = "https://www.lit.edu.cn/jwc/xb2021/"
)

type Post struct {
	Title string
	Url   string
	Date  string
}

type PostList struct {
	Posts []Post
	Next  string
}

// JwUser 教务在线结构体
type JwCUser struct {
	Client *resty.Client
}

// NewJwUser 新建教务用户
func NewJwCUser() *JwCUser {

	var u JwCUser
	u.Client = resty.New()
	return &u
}

// GetGGTZ 获取公告通知
func (u *JwCUser) GetGGTZ(nextPath string) (string, error) {

	var url string

	if nextPath != "" {
		url = GGTZ_PAGE_URL + nextPath
	} else {
		url = GGTZ_URL
	}

	resp, err := u.Client.R().Get(url)
	if err != nil {
		return "", err
	}
	return resp.String(), nil
}

// 获取公告通知
func (u *JwCUser) GetGGTZPost(nextPath string) (list PostList, err error) {

	data, err := u.GetGGTZ(nextPath)
	if err != nil {
		return
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(data))
	if err != nil {
		return
	}

	// log.Println(doc.Html())

	var posts []Post

	doc.Find("ul.list_news").First().Find("li").Each(func(_ int, s *goquery.Selection) {
		// For each item found, get the band and title
		title := s.Find("a").Text()
		url, _ := s.Find("a").Attr("href")
		date := s.Find("span").Text()
		// log.Println(title, url, date)
		posts = append(posts, Post{
			Title: title,
			Url:   url,
			Date:  date,
		})
	})

	next, _ := doc.Find("a.Next").First().Attr("href")

	list.Posts = posts
	list.Next = next

	if next != "" {
		if strings.Contains(next, "ggtz") {
			list.Next = next
		} else {
			list.Next = "ggtz/" + next
		}
	}

	// log.Println(

	return
}
