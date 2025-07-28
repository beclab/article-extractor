package templates

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func theMirrorScrapContent(document *goquery.Document) string {
	contents := ""
	document.Find("div.factbox").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})
	document.Find("p").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		if strings.HasPrefix(text, "READ MORE:") {
			RemoveNodes(s)
		}

	})
	document.Find("div.article-body").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) TheMirrorExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := theMirrorScrapContent(document)
	return content, "", 0, "", "", ""
}
