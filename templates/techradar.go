package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func techradarScrapContent(document *goquery.Document) string {
	contents := ""

	document.Find("aside,h3#section-you-might-also-like,ul,div.slice-container").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})

	document.Find("div.hero-image-wrapper,div#article-body").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) TechradarExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := techradarScrapContent(document)
	return content, "", 0, "", "", ""
}
