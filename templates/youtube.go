package templates

import (
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

func (t *Template) YoutubeScrapContent(document *goquery.Document) string {
	contents := ""

	/*document.Find("#description-inline-expander").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})*/
	return contents
}

func (t *Template) YoutubeMediaContent(url string, document *goquery.Document) (string, string, string) {
	//pattern := `^https?://(?:www\.)?youtube\.com/watch\?v=([a-zA-Z0-9_-]+).*$`
	pattern := `youtube\.com/watch\?v=([^&]+)`
	regex := regexp.MustCompile(pattern)
	match := regex.FindStringSubmatch(url)
	if match != nil {
		if len(match) > 1 {
			videoID := match[1]
			embedUrl := "https://www.youtube.com/embed/gfx7mTmWdYU?si=" + videoID
			contents := "<iframe width='840' height='472' src='" + embedUrl + "'  frameborder='0' allow='accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share' referrerpolicy='strict-origin-when-cross-origin' allowfullscreen></iframe>"
			return contents, url, "video"
		}

	}
	return "", "", ""

}
