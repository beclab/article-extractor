package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func nprScrapContent(document *goquery.Document) string {
	contents := ""

	document.Find("div.internallink,div.enlarge-options,div.enlarge_measure,div.enlarge_html").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)

	})
	document.Find("div#storytext").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) NprExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := nprScrapContent(document)
	return content, "", 0, "", "", ""
}
