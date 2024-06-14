package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func (t *Template) TechspotScrapContent(document *goquery.Document) string {
	contents := ""

	document.Find("header#content-header,div.feature-title,aside,ul.social-widget,div.related-news,div.comment-count-block,div.social-share-svg,div.related-products,section.related-query,section.related-products,nav,div.subDriveRevBot").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})

	document.Find("article").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}
