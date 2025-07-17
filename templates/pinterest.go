package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func pinterestScrapContent(document *goquery.Document) string {
	contents := ""
	document.Find("div[data-test-id=pin-closeup-image],div[data-test-id=main-pin-description-text],div[data-test-id=truncated-description]").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) PinterestExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := pinterestScrapContent(document)
	author := ""
	embedUrl := ""
	document.Find("div[data-test-id=creator-profile-name]").Each(func(i int, s *goquery.Selection) {
		author = s.Text()
	})
	document.Find("video").Each(func(i int, s *goquery.Selection) {
		embedUrl = url
	})

	return content, author, 0, embedUrl, embedUrl, ""
}
