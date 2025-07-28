package templates

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func timeContentExtractor(document *goquery.Document) string {
	contents := ""
	document.Find("p").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		if strings.HasPrefix(text, "Read More:") {
			RemoveNodes(s)
		}
	})

	document.Find("div.featured-media,div#article-body-main").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) TimeExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := timeContentExtractor(document)
	return content, "", 0, "", "", ""
}
