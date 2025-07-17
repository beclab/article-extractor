package templates

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ewScrapContent(document *goquery.Document) string {

	contents := ""
	document.Find("div.article__primary-video-container").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})

	document.Find("p").Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		if strings.HasPrefix(text, "Sign up") {
			RemoveNodes(s)
		}

	})

	document.Find("div.article-content>div.article-content").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) EWExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := ewScrapContent(document)

	return content, "", 0, "", "", ""
}
