package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func fandomScrapContent(document *goquery.Document) string {
	contents := ""

	document.Find("div#content").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})

	return contents
}

func (t *Template) FandomExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := fandomScrapContent(document)

	return content, "", 0, "", "", ""
}
