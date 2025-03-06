package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func (t *Template) RedditScrapContent(document *goquery.Document) string {
	contents := ""

	document.Find("div[slot='text-body']").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}
