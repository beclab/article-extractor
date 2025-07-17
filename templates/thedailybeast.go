package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func thedailybeastScrapContent(document *goquery.Document) string {

	contents := ""
	document.Find("section.Hero,article.Body").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) ThedailybeastExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := thedailybeastScrapContent(document)
	return content, "", 0, "", "", ""
}
