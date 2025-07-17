package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func geektyrantScrapContent(document *goquery.Document) string {

	contents := ""

	document.Find("div.sqs-block-content").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) GeektyrantExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := geektyrantScrapContent(document)

	return content, "", 0, "", "", ""
}
