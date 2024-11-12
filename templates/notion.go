package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func (t *Template) NotionScrapContent(document *goquery.Document) string {
	contents := ""

	/*document.Find("aside,div.o-ads").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})*/
	document.Find("div.notion-page-content").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}
