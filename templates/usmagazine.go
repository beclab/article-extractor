package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func (t *Template) UsmagazineScrapContent(document *goquery.Document) string {
	contents := ""

	document.Find("div#news-block,div.link-related,div.in-this-article").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)

	})
	document.Find("div.article-content").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}
