package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func bbcScrapContent(document *goquery.Document) string {
	contents := ""
	document.Find("header,div[data-testid=byline],div[data-component=topic-list],div[data-component=links-block],span.visually-hidden,h1#main-heading,div[data-component=byline-block],div[data-component=timestamp-block],div[data-component=headline-block],div[data-component=tags]").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})
	document.Find("div.description").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents = content
	})
	if contents == "" {
		document.Find("article,#main-content").Each(func(i int, s *goquery.Selection) {
			var content string
			content, _ = goquery.OuterHtml(s)
			contents = content
		})
	}
	return contents
}

func (t *Template) BBCExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := bbcScrapContent(document)
	return content, "", 0, "", "", ""
}
