package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func (t *Template) ThewrapScrapContent(document *goquery.Document) string {

	contents := ""
	document.Find("div.wp-block-the-wrap-ad,div.wp-block-the-wrap-read-more,figure.wp-block-embed").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})
	document.Find("figure.wp-block-post-featured-image>img,div.entry-content").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}
