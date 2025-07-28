package templates

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func indiatimesScrapContent(document *goquery.Document) string {
	contents := ""

	document.Find("header,div.outsideInd").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})
	document.Find("strong").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		if strings.HasPrefix(text, "Also Read") {
			RemoveNodes(s)
		}
	})
	document.Find("figure.artImg,div.artText").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) IndiatimesExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := indiatimesScrapContent(document)
	return content, "", 0, "", "", ""
}
