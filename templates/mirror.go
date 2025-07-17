package templates

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func mirrorScrapContent(document *goquery.Document) string {
	contents := ""
	document.Find("aside,div#mantis-recommender-wrapper").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})
	document.Find("p").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		if strings.HasPrefix(text, "READ MORE:") || strings.HasPrefix(text, "REA MORE:") {
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

func (t *Template) MirrorExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := mirrorScrapContent(document)
	return content, "", 0, "", "", ""
}
