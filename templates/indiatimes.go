package templates

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func (t *Template) IndiatimesScrapContent(document *goquery.Document) string {
	contents := ""

	document.Find("header,a,div.outsideInd").Each(func(i int, s *goquery.Selection) {
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
