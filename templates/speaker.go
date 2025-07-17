package templates

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func spreakerScrapContent(document *goquery.Document) string {
	contents := ""

	document.Find("div[x-show=collapsed]").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})
	document.Find("div.text-md").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) SpreakerExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := spreakerScrapContent(document)
	audioUrl := ""
	document.Find("meta[name='twitter:player']").Each(func(i int, s *goquery.Selection) {
		content, _ := s.Attr("content")
		//"https://widget.spreaker.com/player?episode_id=60819052&playlist=show&cover_image_url=https%3A%2F%2Fd3wo5wojvuv7l.cloudfront.net%2Fimages.spreaker.com%2Foriginal%2F551ce348940065825b0a755b58fdb5ae.jpg"
		startIndex := strings.Index(content, "episode_id=") + len("episode_id=")
		endIndex := strings.Index(content[startIndex:], "&")
		if startIndex > -1 && endIndex > -1 {
			episodeID := content[startIndex : startIndex+endIndex]
			audioUrl = "https://api.spreaker.com/v2/episodes/" + episodeID + "/ondemand.mp3"
		}
	})
	return content, "", 0, audioUrl, audioUrl, "audio"
}
