package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func eonlineScrapContent(document *goquery.Document) string {
	contents := ""
	document.Find("div.article-detail__main-content").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) EOnlineExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := eonlineScrapContent(document)

	return content, "", 0, "", "", ""
}
