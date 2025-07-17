package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func businessLiveScrapContent(document *goquery.Document) string {

	contents := ""
	document.Find("div.article-widget-related_articles").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})

	document.Find("div.article-widgets").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents

}

func (t *Template) BusinessLiveExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := businessLiveScrapContent(document)
	return content, "", 0, "", "", ""
}
