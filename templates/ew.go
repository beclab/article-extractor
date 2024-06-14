package templates

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func (t *Template) EWScrapContent(document *goquery.Document) string {

	contents := ""
	document.Find("div.article__primary-video-container").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})

	document.Find("p").Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		if strings.HasPrefix(text, "Sign up") {
			RemoveNodes(s)
		}

	})

	document.Find("div.article-content>div.article-content").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}
