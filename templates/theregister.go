package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func theRegisterScrapContent(document *goquery.Document) string {

	contents := ""
	document.Find("ul.listinks,div[aria-hidden=true]").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})
	document.Find("div#body").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) TheRegisterExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := theRegisterScrapContent(document)
	return content, "", 0, "", "", ""
}
