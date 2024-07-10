package templates

import (
	"encoding/json"
	"log"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func (t *Template) VimeoScrapContent(document *goquery.Document) string {
	contents := ""
	scriptSelector := "script[type=\"application/ld+json\"]"
	document.Find(scriptSelector).Each(func(i int, s *goquery.Selection) {
		scriptContent := strings.TrimSpace(s.Text())
		var metaData []map[string]interface{}
		unmarshalErr := json.Unmarshal([]byte(scriptContent), &metaData)
		if unmarshalErr != nil || len(metaData) < 1 {
			log.Printf("convert  unmarshalError %v", unmarshalErr)
		}
		if description, ok := metaData[0]["description"]; ok {
			contents = description.(string)
		}

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
