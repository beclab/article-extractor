package templates

import (
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

func (t *Template) VimeoScrapContent(document *goquery.Document) string {
	contents := ""
	document.Find("div.clip_details-description").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) VimeoMediaContent(url string, document *goquery.Document) (string, string, string) {
	pattern := `vimeo\.com/(\d+)`
	regex := regexp.MustCompile(pattern)
	match := regex.FindStringSubmatch(url)

	if match != nil {
		if len(match) > 1 {
			videoID := match[1]
			embedUrl := "https://player.vimeo.com/video/" + videoID
			contents := "<iframe width='896' height='504' src='" + embedUrl + "' frameborder='0' referrerpolicy='no-referrer'></iframe>"
			return contents, url, "video"
		}

	}
	return "", "", ""
}
