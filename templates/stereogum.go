package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func stereogumScrapContent(document *goquery.Document) string {

	contents := ""

	document.Find("h1.headline,div.article-meta-information,div.article-author,div.article-tags").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})
	document.Find("div.article__content").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) StereogumExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := stereogumScrapContent(document)
	return content, "", 0, "", "", ""
}
