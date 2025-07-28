package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func digitalTrendsScrapContent(document *goquery.Document) string {

	contents := ""
	document.Find("div.b-related-links,ul.h-editors-recs,h4.h-editors-recs-title,div#dt-toc").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})

	document.Find("article").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) DigitalTrendsExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := digitalTrendsScrapContent(document)

	return content, "", 0, "", "", ""
}
