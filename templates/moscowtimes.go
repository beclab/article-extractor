package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func moscowTimesScrapContent(document *goquery.Document) string {
	contents := ""
	document.Find("aside").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})
	document.Find("figure.article__featured-image,div.article__content").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) MoscowTimesExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := moscowTimesScrapContent(document)
	return content, "", 0, "", "", ""
}
