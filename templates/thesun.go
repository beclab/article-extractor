package templates

import "github.com/PuerkitoBio/goquery"

func (t *Template) TheSunScrapContent(document *goquery.Document) string {
	contents := ""
	document.Find("div.read-more-container,div.advert-wrapper,div.rail--classic,div.article__gallery-count").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})
	document.Find("div.article__content").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}
