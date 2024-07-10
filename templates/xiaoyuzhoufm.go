package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func (t *Template) XiaoyuzhouFMScrapContent(document *goquery.Document) string {
	contents := ""

	document.Find("div.sn-content").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) XiaoyuzhouFMMediaContent(url string, document *goquery.Document) (string, string, string) {
	audioUrl := ""
	document.Find("meta[property='og:audio']").Each(func(i int, s *goquery.Selection) {
		if content, exists := s.Attr("content"); exists {
			audioUrl = content
		}

	})
	return audioUrl, audioUrl, "audio"
}
