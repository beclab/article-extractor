package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func pagesixScrapContent(document *goquery.Document) string {
	contents := ""
	document.Find("aside,div.comments-inline-cta,div.inline-module").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})

	document.Find("div.featured-image,div.entry-content").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) PagesixExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := pagesixScrapContent(document)
	return content, "", 0, "", "", ""
}
