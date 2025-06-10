package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func (t *Template) FandomScrapContent(document *goquery.Document) string {
	contents := ""

	document.Find("div#content").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})

	return contents
}
