package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func (t *Template) BenzingaScrapContent(document *goquery.Document) string {
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
