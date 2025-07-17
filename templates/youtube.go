package templates

import (
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

func (t *Template) YoutubeExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	pattern := `youtube\.com/watch\?v=([^&]+)`
	regex := regexp.MustCompile(pattern)
	match := regex.FindStringSubmatch(url)
	if match != nil {
		if len(match) > 1 {
			videoID := match[1]
			embedUrl := "https://www.youtube.com/embed/gfx7mTmWdYU?si=" + videoID
			contents := "<iframe width='840' height='472' src='" + embedUrl + "'  frameborder='0' allow='accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share' referrerpolicy='strict-origin-when-cross-origin' allowfullscreen></iframe>"
			return "", "", 0, contents, url, "video"
		}

	}
	return "", "", 0, "", "", ""
}
