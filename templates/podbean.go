package templates

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func (t *Template) PodBeanScrapContent(document *goquery.Document) string {
	contents := ""

	document.Find("div.episode-detail-top,div.cc-post-toolbar").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})
	document.Find("div.episode-content").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) PodBeanMediaContent(url string, document *goquery.Document) (string, string, string) {
	audioUrl := ""
	scriptSelector := "script[type=\"application/ld+json\"]"
	document.Find(scriptSelector).Each(func(i int, s *goquery.Selection) {
		scriptContent := strings.TrimSpace(s.Text())
		var metaData map[string]interface{}
		unmarshalErr := json.Unmarshal([]byte(scriptContent), &metaData)
		if unmarshalErr != nil {
			log.Printf("convert  unmarshalError %v", unmarshalErr)
		}
		if associatedMediaData, ok := metaData["associatedMedia"]; ok {
			switch associatedMediaData.(type) {
			case map[string]interface{}:
				associatedMediaDetail := associatedMediaData.(map[string]interface{})
				if contentUrl, ok := associatedMediaDetail["contentUrl"]; ok {
					audioUrl = contentUrl.(string)
				}
			}
		}

	})
	return audioUrl, audioUrl, "audio"
}
