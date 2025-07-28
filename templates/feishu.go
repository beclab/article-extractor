package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func feishuScrapContent(document *goquery.Document) string {
	contents := ""

	document.Find("div.page-block-children").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) FeishuExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := feishuScrapContent(document)

	return content, "", 0, "", "", ""
}
