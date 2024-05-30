package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func (t *Template) VoxScrapContent(document *goquery.Document) string {
	contents := ""

	document.Find("div.duet--layout--rail,aside,div._1agbrixh").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)

	})
	document.Find("div.duet--layout--entry-body-container").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}
