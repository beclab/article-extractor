package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func (t *Template) MITScrapContent(document *goquery.Document) string {
	contents := ""

	document.Find("div.news-article--image-item,div.news-article--content--body--inner").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}
