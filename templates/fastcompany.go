package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func fastcompanyScrapContent(document *goquery.Document) string {
	contents := ""

	document.Find("div.ad-container").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)

	})
	document.Find("header>img,article.article-container").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) FastcompanyExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := fastcompanyScrapContent(document)

	return content, "", 0, "", "", ""
}
