package templates

import (
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

func yahooContentExtractor(document *goquery.Document) string {
	contents := ""
	document.Find("header.caas-header,div.caas-content-byline-wrapper,button,div.xray-error-wrapper,div.caas-xray-pills-container,aside.caas-aside-section").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})
	document.Find("div#module-article").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) YahooExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := yahooContentExtractor(document)
	author := ""
	document.Find("span.caas-author-byline-collapse").Each(func(i int, s *goquery.Selection) {
		reg := regexp.MustCompile(`(?:\n\s+)`)
		author = reg.ReplaceAllString(s.Text(), "")

	})
	return content, author, 0, "", "", ""
}
