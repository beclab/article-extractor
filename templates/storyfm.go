package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func (t *Template) StoryFMScrapContent(document *goquery.Document) string {
	contents := ""

	document.Find("div.rs-post__content").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) StoryFMMediaContent(url string, document *goquery.Document) (string, string, string) {
	audioUrl := ""
	document.Find("audio.sf-audio > source").Each(func(i int, s *goquery.Selection) {
		audioUrl, _ = s.Attr("src")
	})
	return audioUrl, audioUrl, "audio"
}
