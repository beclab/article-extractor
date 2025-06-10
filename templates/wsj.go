package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func (t *Template) WsjScrapContent(document *goquery.Document) string {
	contents := ""

	document.Find("div[data-testid=ad-container]").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})
	document.Find("div.meteredContent").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}
