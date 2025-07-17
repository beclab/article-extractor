package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func pitchForkScrapContent(document *goquery.Document) string {
	contents := ""
	document.Find("figure.iframe-embed").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})
	document.Find("div.lead-asset,div.body__inner-container").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) PitchForkExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := pitchForkScrapContent(document)
	return content, "", 0, "", "", ""
}
