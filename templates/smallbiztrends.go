package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func smallBizTrendsScrapContent(document *goquery.Document) string {

	contents := ""
	document.Find("h1.entry-title,span.byline").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})

	document.Find("div.post-inner").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents

}

func (t *Template) SmallBizTrendsExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := smallBizTrendsScrapContent(document)
	return content, "", 0, "", "", ""
}
