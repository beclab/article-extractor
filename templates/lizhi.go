package templates

import (
	"html"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func (t *Template) LizhiMediaContent(url string, document *goquery.Document) (string, string, string) {
	audioUrl := ""
	document.Find("script").Each(func(i int, s *goquery.Selection) {
		scriptContent, err := s.Html()
		if audioUrl == "" && err == nil && scriptContent != "" {
			content := html.UnescapeString(strings.ReplaceAll(scriptContent, "\\", ""))

			re := regexp.MustCompile(`\"voiceTrack\":\"(http[^"]+)"`)
			matches := re.FindStringSubmatch(content)

			if len(matches) > 1 {
				audioUrl = matches[1]
				return
			}
		}
	})

	return audioUrl, audioUrl, "audio"
}
