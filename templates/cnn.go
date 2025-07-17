package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func cnnScrapContent(document *goquery.Document) string {

	contents := ""

	document.Find("main.article__main").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) CNNExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := cnnScrapContent(document)

	return content, "", 0, "", "", ""
}
