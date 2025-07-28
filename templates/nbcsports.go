package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func nbcSportsScrapContent(document *goquery.Document) string {
	contents := ""
	document.Find("div.ArticlePage-articleBody").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) NBCSportsExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := nbcSportsScrapContent(document)
	return content, "", 0, "", "", ""
}
