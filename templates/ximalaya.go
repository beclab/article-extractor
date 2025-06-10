package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func (t *Template) XimalayaScrapContent(document *goquery.Document) string {
	contents := ""

	document.Find("article.intro").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) XimalayaScrapMetaData(document *goquery.Document) (string, string) {

	author := ""
	published_at := ""

	document.Find("a.albumTitle").Each(func(i int, s *goquery.Selection) {
		author = s.Text()
	})
	return author, published_at
}

func (t *Template) XimalayaMediaContent(url string, document *goquery.Document) (string, string, string) {
	return url, url, "audio"
}
