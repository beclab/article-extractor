package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func storyFMScrapContent(document *goquery.Document) string {
	contents := ""

	document.Find("div.rs-post__content").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) StoryFMExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := storyFMScrapContent(document)
	audioUrl := ""
	fileType := ""
	document.Find("audio.sf-audio > source").Each(func(i int, s *goquery.Selection) {
		audioUrl, _ = s.Attr("src")
		fileType = AudioFileType
	})
	return content, "", 0, audioUrl, audioUrl, fileType
}
