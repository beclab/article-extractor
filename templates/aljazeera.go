package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func aljazeeraScrapContent(document *goquery.Document) string {
	contents := ""
	document.Find("div.more-on").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})

	document.Find("figure.article-featured-image,figure.gallery-featured-image,div.wysiwyg--all-content,figure.gallery-image").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) AljazeeraExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := aljazeeraScrapContent(document)
	return content, "", 0, "", "", ""
}
