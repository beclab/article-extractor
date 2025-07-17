package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func dwScrapContent(document *goquery.Document) string {
	contents := ""
	document.Find("header>div,header>h1,header>span,section[data-tracking-name=sharing-icons-inline],div.advertisement,footer").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})

	document.Find("article").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) DWExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := dwScrapContent(document)

	return content, "", 0, "", "", ""
}
