package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func (t *Template) SubStackScrapContent(document *goquery.Document) string {
	contents := ""

	document.Find("div[class^='subscription-widget'],p.button-wrapper").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})

	document.Find("div.available-content").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}
