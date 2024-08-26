package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func (t *Template) PinterestScrapContent(document *goquery.Document) string {
	contents := ""

	document.Find("div[data-test-id=pin-closeup-image],div[data-test-id=main-pin-description-text],div[data-test-id=truncated-description]").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) PinterestScrapMetaData(document *goquery.Document) (string, string) {

	author := ""
	published_at := ""
	document.Find("div[data-test-id=creator-profile-name]").Each(func(i int, s *goquery.Selection) {
		author = s.Text()
	})

	return author, published_at
}

func (t *Template) PinterestMediaContent(url string, document *goquery.Document) (string, string, string) {
	contents := ""
	embedUrl := ""
	mediaType := ""
	document.Find("video").Each(func(i int, s *goquery.Selection) {
		embedUrl = url
		mediaType = "video"
	})
	return contents, embedUrl, mediaType

}
