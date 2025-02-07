package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func (t *Template) ZhihuScrapContent(document *goquery.Document) string {
	contents := ""

	document.Find("div.Post-RichTextContainer").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) ZhihuScrapMetaData(document *goquery.Document) (string, string) {

	author := ""
	published_at := ""
	document.Find("span.AuthorInfo-name>a").Each(func(i int, s *goquery.Selection) {
		author = s.Text()
	})
	return author, published_at
}
