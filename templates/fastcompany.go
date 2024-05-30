package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func (t *Template) FastcompanyScrapContent(document *goquery.Document) string {
	contents := ""

	document.Find("div.ad-container").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)

	})
	document.Find("header>img,article.article-container").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}
