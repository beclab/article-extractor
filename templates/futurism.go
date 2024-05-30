package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func (t *Template) FuturismScrapContent(document *goquery.Document) string {
	contents := ""

	document.Find("div.ad__align,aside,img[aria-hidden=true]").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)

	})

	document.Find("article>header").Each(func(i int, s *goquery.Selection) {
		var content string
		next := s.Next()
		if next.Length() > 0 {
			content, _ = goquery.OuterHtml(next)
			contents += content
		}
	})
	document.Find("section#incArticle,div#incArticle").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}
