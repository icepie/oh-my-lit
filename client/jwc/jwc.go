package jwc

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"

	md "github.com/JohannesKaufmann/html-to-markdown"
)

const (
	// ggtz
	GGTZ_URL = "https://www.lit.edu.cn/jwc/xb2021/ggtz.htm"
	// gttz by page
	GGTZ_PAGE_URL = "https://www.lit.edu.cn/jwc/xb2021"
	// post url
	POST_URL = "https://www.lit.edu.cn/jwc"
)

type Attachment struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}

type PostContent struct {
	Content    string       `json:"content"`
	Attachment []Attachment `json:"attachment"`
}

type Post struct {
	Title string `json:"title"`
	Url   string `json:"url"`
	Date  string `json:"date"`
}

type PostList struct {
	Posts []Post `json:"posts"`
	Next  string `json:"next"`
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

func (u *JwCUser) GetGGTZPost(postPath string) (content PostContent, err error) {

	content.Attachment = make([]Attachment, 0)

	data, err := u.getGGTZPostPage(postPath)
	if err != nil {
		return
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(data))
	if err != nil {
		return
	}

	// 获取附件

	// <ul style="list-style-type:none;">
	//     <li>附件【<a href="/system/_content/download.jsp?urltype=news.DownloadAttachUrl&amp;owner=1404874179&amp;wbfileid=9607F8947E5547D2CA3A541D32E985A6&amp;deal=1" target="_blank">小学期计划安排表.xls</a>】已下载<span id="nattach2278921"><script language="javascript">getClickTimes(2278921,1404874179,"wbnewsfile","attach")</script></span>次</li>
	// </ul>

	// 标题

	title := doc.Find("h1").Text()

	html, err := doc.Find("div.v_news_content").Html()
	if err != nil {
		return
	}

	converter := md.NewConverter("", true, nil)

	markdown, err := converter.ConvertString(html)
	if err != nil {
		log.Fatal(err)
	}

	// 替换/*.png 为 https://www.lit.edu.cn/*.png

	// reg := regexp.MustCompile(`\!\[.*\]\((.*)\)`)
	// markdown = reg.ReplaceAllStringFunc(markdown, func(s string) string {
	// 	return strings.Replace(s, "(", "(https://www.lit.edu.cn", 1)
	// })

	// /__local/ to https://www.lit.edu.cn/__local/

	markdown = strings.Replace(markdown, "/__local/", "https://www.lit.edu.cn/__local/", -1)

	// 多个换行替换为两个换行
	reg := regexp.MustCompile(`\n{2,}`)
	markdown = reg.ReplaceAllString(markdown, "\n\n")

	// 替换一行中的多个空格为一个空格
	reg = regexp.MustCompile(` {2,}`)
	markdown = reg.ReplaceAllString(markdown, "")

	content.Content = markdown

	doc.Find("a[target=_blank]").Each(func(_ int, s *goquery.Selection) {
		// For each item found, get the band and title
		url, _ := s.Attr("href")
		name := s.Text()
		if strings.Contains(url, "download.jsp") {

			// /system/ to https://www.lit.edu.cn/system/

			url = strings.Replace(url, "/system/", "https://www.lit.edu.cn/system/", 1)

			content.Attachment = append(content.Attachment, Attachment{
				Url:  url,
				Name: name,
			})
		}
	})

	if title != "" {
		content.Content = fmt.Sprintf("# %s\n\n", title) + content.Content
	}

	if len(content.Attachment) > 0 {

		// 在md 最后添加附件列表

		// 分割线

		content.Content += "\n\n---\n\n"

		content.Content += "\n\n**附件列表**\n\n"

		for _, v := range content.Attachment {
			content.Content += fmt.Sprintf("* [%s](%s)\n", v.Name, v.Url)
		}

	}

	// var content string

	// doc.Find("div#vsb_content").First().Find("p").Each(func(_ int, s *goquery.Selection) {
	// 	// For each item found, get the band and title
	// 	content += s.Text()
	// })

	return
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
