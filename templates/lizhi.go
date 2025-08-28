package templates

import (
	"html"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func (t *Template) LizhiExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	audioUrl := ""
	fileType := ""
	document.Find("script").Each(func(i int, s *goquery.Selection) {
		scriptContent, err := s.Html()
		if audioUrl == "" && err == nil && scriptContent != "" {
			content := html.UnescapeString(strings.ReplaceAll(scriptContent, "\\", ""))

			re := regexp.MustCompile(`\"voiceTrack\":\"(http[^"]+)"`)
			matches := re.FindStringSubmatch(content)

			if len(matches) > 1 {
				audioUrl = matches[1]
				fileType = AudioFileType
				return
			}
		}
	})
	return "", "", 0, audioUrl, audioUrl, fileType
}
