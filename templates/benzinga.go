package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func benzingaScrapContent(document *goquery.Document) string {
	contents := ""
	document.Find("div.bz-campaign,div.animate-pulse,img[aria-hidden=true]").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)

	})
	document.Find("div.key-points-wrapper,div#article-body").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) BenzingaExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := benzingaScrapContent(document)
	return content, "", 0, "", "", ""
}
