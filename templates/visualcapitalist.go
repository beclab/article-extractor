package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func (t *Template) VisualcapitalistScrapContent(document *goquery.Document) string {
	contents := ""

	document.Find("div.mvp-author-info-wrap,div.ss-inline-share-wrapper,div.vce-yml-block,div.vce-sponsor-disclaimer,div.vc-info-box,a.vc-newsletter,div.wpforms-container").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})
	document.Find("div.vc-full-width-wrapper").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}
