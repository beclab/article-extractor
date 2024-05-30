package templates

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func (t *Template) TheMirrorScrapContent(document *goquery.Document) string {
	contents := ""
	document.Find("div.factbox").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})
	document.Find("p").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		if strings.HasPrefix(text, "READ MORE:") {
			RemoveNodes(s)
		}

	})
	document.Find("div.article-body").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}
