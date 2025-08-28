package templates

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func bilibiliScrapContent(document *goquery.Document) string {
	contents := ""
	document.Find("span.desc-info-text,div.opus-module-content").Each(func(i int, s *goquery.Selection) {
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
				content = strings.TrimPrefix(content, "https:")
			}
			add_img = fmt.Sprintf("<figure><img src=\"%s\"/></figure>", content)
		}
	})
	return add_img + contents
}

func (t *Template) BilibiliExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	author := ""
	document.Find("div.fixed-author-header__author__name,div.opus-module-author__name").Each(func(i int, s *goquery.Selection) {
		author = strings.TrimSpace(s.Text())
	})
	var publishedAt int64 = 0
	document.Find("div.opus-module-author__pub__text").Each(func(i int, s *goquery.Selection) {
		publishTimes := s.Text()
		layout := "2006年01月02日 15:04"
		publishTimes = strings.TrimPrefix(publishTimes, "编辑于 ")
		publishedAt, _ = ParseLocationTimestamp(publishTimes, layout, ShanghaiTZ)
	})
	content := bilibiliScrapContent(document)

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
		return content, author, publishedAt, contents, url, VideoFileType
	}
	if strings.Contains(url, "bilibili.com/festival/") {
		return content, author, publishedAt, "", url, VideoFileType
	}
	if strings.Contains(url, "audio/au") {
		return content, author, publishedAt, "", url, AudioFileType
	}

	return content, author, publishedAt, "", "", ""
}
