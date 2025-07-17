package templates

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func hbrScrapContent(document *goquery.Document) string {
	contents := ""
	document.Find("div.article-body").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) HBRExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := hbrScrapContent(document)
	author := ""
	document.Find("content[js-target='article-content']").Each(func(i int, s *goquery.Selection) {
		pText := s.Find("p").First().Text()
		if strings.HasPrefix(pText, "By ") {
			author = strings.TrimPrefix(pText, "By ")
		}
	})
	return content, author, 0, "", "", ""
}
