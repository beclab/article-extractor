package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func mitScrapContent(document *goquery.Document) string {
	contents := ""

	document.Find("div.news-article--image-item,div.news-article--content--body--inner").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) MITExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := mitScrapContent(document)
	return content, "", 0, "", "", ""
}
