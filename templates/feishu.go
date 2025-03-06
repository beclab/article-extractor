package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func (t *Template) FeishuScrapContent(document *goquery.Document) string {
	contents := ""

	document.Find("div.page-block-children").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}
