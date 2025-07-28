package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func (t *Template) WsjExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	contents := ""
	document.Find("div[data-testid=ad-container]").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})
	document.Find("div.meteredContent").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})

	return contents, "", 0, "", "", ""
}
