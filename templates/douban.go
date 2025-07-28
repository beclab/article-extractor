package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func doubanScrapContent(document *goquery.Document) string {

	contents := ""
	document.Find("div.topic-richtext,div.review-content").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) DoubanExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := doubanScrapContent(document)

	return content, "", 0, "", "", ""
}
