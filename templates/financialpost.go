package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func (t *Template) FinancialPostScrapContent(document *goquery.Document) string {
	contents := ""

	document.Find("div.featured-video,div.visually-hidden,h2.visually-hidden").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)

	})
	document.Find("section.article-content__content-group").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}
