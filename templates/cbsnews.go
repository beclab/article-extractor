package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func cbsNewsScrapContent(document *goquery.Document) string {
	contents := ""
	document.Find("aside,div.content-author,ul.content__tags").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})
	document.Find("section.content__body").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) CbsNewsExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := cbsNewsScrapContent(document)
	return content, "", 0, "", "", ""
}
