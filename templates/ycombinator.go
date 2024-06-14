package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func (t *Template) YcombinatorScrapContent(document *goquery.Document) string {
	contents := ""

	/*document.Find("header").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)

	})*/
	document.Find("div.prose").Each(func(i int, s *goquery.Selection) {
		var content string
		prev := s.Prev()
		if prev.Length() > 0 {
			content, _ = goquery.OuterHtml(prev)
			contents += content
		}

		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}
