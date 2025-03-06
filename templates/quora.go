package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func (t *Template) QuoraScrapContent(document *goquery.Document) string {
	contents := ""

	document.Find("div.puppeteer_test_answer_content']").Each(func(i int, s *goquery.Selection) {
		parent := s.Parent()
		var content string
		content, _ = goquery.OuterHtml(parent)
		contents += content
	})
	return contents
}
