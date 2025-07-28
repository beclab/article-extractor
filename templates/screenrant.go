package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func screenrantScrapContent(document *goquery.Document) string {
	contents := ""

	document.Find("div.adsninja-ad-zone,div.active-content").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})
	document.Find("div.heading_image,section.article-body").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) ScreenrantExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := screenrantScrapContent(document)
	return content, "", 0, "", "", ""
}
