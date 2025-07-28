package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func thevergeScrapContent(document *goquery.Document) string {

	contents := ""

	document.Find("button,div.duet--recirculation--related-list").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})
	document.Find("figure.duet--article--lede-image,div.duet--article--article-body-component-container").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) ThevergeExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := thevergeScrapContent(document)
	return content, "", 0, "", "", ""
}
