package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func ftScrapContent(document *goquery.Document) string {
	contents := ""

	document.Find("aside,div.o-ads").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})
	document.Find("div.main-image,article#article-body").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) FTExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := ftScrapContent(document)

	return content, "", 0, "", "", ""
}
