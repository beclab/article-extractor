package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func deadspinScrapContent(document *goquery.Document) string {

	contents := ""

	document.Find("div.js_post-content").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) DeadspinExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := deadspinScrapContent(document)

	return content, "", 0, "", "", ""
}
