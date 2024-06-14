package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func (t *Template) ScreencrushScrapContent(document *goquery.Document) string {
	contents := ""

	document.Find("div.bands-in-town-container,div.tsm-ad,div.branded-app-shortcode-inarticle,iframe").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})
	document.Find("div.content").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}
