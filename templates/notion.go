package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func notionScrapContent(document *goquery.Document) string {
	contents := ""
	document.Find("div.page-body").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) NotionExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := notionScrapContent(document)
	return content, "", 0, "", "", ""
}
