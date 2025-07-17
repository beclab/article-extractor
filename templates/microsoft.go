package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func microsoftScrapContent(document *goquery.Document) string {
	contents := ""

	document.Find("header").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)

	})
	document.Find("article.m-blog-post").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) MicrosoftExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := microsoftScrapContent(document)
	return content, "", 0, "", "", ""
}
