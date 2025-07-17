package templates

import "github.com/PuerkitoBio/goquery"

func theSunScrapContent(document *goquery.Document) string {
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

func (t *Template) TheSunExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := theSunScrapContent(document)
	return content, "", 0, "", "", ""
}
