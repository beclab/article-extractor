package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func ibtimesScrapContent(document *goquery.Document) string {
	contents := ""

	document.Find("header").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)

	})
	document.Find("article.node-article").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) IbtimesExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := ibtimesScrapContent(document)
	return content, "", 0, "", "", ""
}
