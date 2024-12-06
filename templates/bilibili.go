package templates

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func (t *Template) BilibiliScrapContent(document *goquery.Document) string {
	contents := ""

	document.Find("span.desc-info-text").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	if contents != "" {
		return contents
	}
	document.Find("p[class*='mediainfo_content_placeholder']").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	add_img := ""
	document.Find("meta[property='og:image']").Each(func(i int, s *goquery.Selection) {
		if content, exists := s.Attr("content"); exists {
			if strings.HasPrefix(content, "https:https://") {
				// 去掉前面的 "https:"
				content = strings.TrimPrefix(content, "https:")
			}
			add_img = fmt.Sprintf("<figure><img src=\"%s\"/></figure>", content)
		}
	})

	return add_img + contents
}

func (t *Template) BilibiliMediaContent(url string, document *goquery.Document) (string, string, string) {
	bvid := ""
	document.Find("meta[itemprop=url]").Each(func(i int, s *goquery.Selection) {
		if content, exists := s.Attr("content"); exists {
			videoPattern := `video/(\w+)`
			re := regexp.MustCompile(videoPattern)
			match := re.FindStringSubmatch(content)
			if len(match) > 1 {
				bvid = match[1]
			}
		}

	})

	if bvid != "" {
		embeddingUrl := "https://www.bilibili.com/blackboard/html5mobileplayer.html?bvid=" + bvid + "&amp;high_quality=1&amp;autoplay=0"
		contents := "<iframe width='910' height='668' src='" + embeddingUrl + "'  border='0' scrolling='no' border='0 frameborder='no' framespacing='0' allowfullscreen='true' referrerpolicy='no-referrer'></iframe>"
		return contents, url, "video"
	}
	downloadUrl := ""
	downloadType := ""
	document.Find("meta[property='og:url']").Each(func(i int, s *goquery.Selection) {
		if content, exists := s.Attr("content"); exists {
			downloadUrl = content
			downloadType = "video"
		}
	})
	return "", downloadUrl, downloadType
}
