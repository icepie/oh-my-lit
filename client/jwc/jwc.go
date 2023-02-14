package jwc

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"
)

const (
	// ggtz
	GGTZ_URL = "https://www.lit.edu.cn/jwc/xb2021/ggtz.htm"
	// gttz by page
	GGTZ_PAGE_URL = "https://www.lit.edu.cn/jwc/xb2021"
	// post url
	POST_URL = "https://www.lit.edu.cn/jwc"
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

// GetGGTZPage 获取公告通知
func (u *JwCUser) getGGTZPage(nextPath string) (string, error) {

	var url string

	if nextPath != "" {
		url = GGTZ_PAGE_URL + "/" + nextPath
	} else {
		url = GGTZ_URL
	}

	resp, err := u.Client.R().Get(url)
	if err != nil {
		return "", err
	}
	return resp.String(), nil
}

// getPost 获取公告通知
func (u *JwCUser) getGGTZPostPage(postPath string) (string, error) {

	url := POST_URL + "/" + postPath

	resp, err := u.Client.R().Get(url)
	if err != nil {
		return "", err
	}
	return resp.String(), nil
}

//

func (u *JwCUser) GetGGTZPost(postPath string) (string, error) {

	data, err := u.getGGTZPostPage(postPath)
	if err != nil {
		return "", err
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(data))
	if err != nil {
		return "", err
	}

	log.Println(doc.Html())

	// var content string

	// doc.Find("div#vsb_content").First().Find("p").Each(func(_ int, s *goquery.Selection) {
	// 	// For each item found, get the band and title
	// 	content += s.Text()
	// })

	return "", nil
}

// 获取公告通知
func (u *JwCUser) GetGGTZPostList(nextPath string) (list PostList, err error) {

	data, err := u.getGGTZPage(nextPath)
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

		url = strings.Replace(url, "../", "", -1)
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
