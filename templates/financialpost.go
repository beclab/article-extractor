package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func financialPostScrapContent(document *goquery.Document) string {
	contents := ""

	document.Find("div.featured-video,div.visually-hidden,h2.visually-hidden").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)

	})
	document.Find("section.article-content__content-group").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) FinancialPostExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := financialPostScrapContent(document)

	return content, "", 0, "", "", ""
}
