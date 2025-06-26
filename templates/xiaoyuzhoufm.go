package templates

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func (t *Template) XiaoyuzhouFMScrapContent(document *goquery.Document) string {
	contents := ""

	document.Find("div.sn-content").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) XiaoyuzhouScrapMetaData(document *goquery.Document) (string, string) {
	author := ""
	published_at := ""

	scriptSelector := "script[type=\"application/ld+json\"]"
	document.Find(scriptSelector).Each(func(i int, s *goquery.Selection) {
		scriptContent := strings.TrimSpace(s.Text())
		var metaData map[string]interface{}
		unmarshalErr := json.Unmarshal([]byte(scriptContent), &metaData)
		if unmarshalErr != nil {
			log.Printf("convert  unmarshalError %v", unmarshalErr)
		}
		if partOfSeriesData, ok := metaData["partOfSeries"]; ok {
			switch partOfSeriesData.(type) {
			case map[string]interface{}:
				associatedMediaDetail := partOfSeriesData.(map[string]interface{})
				if name, ok := associatedMediaDetail["name"]; ok {
					author = name.(string)
					return
				}
			}
		}

	})
	return author, published_at
}

func (t *Template) XiaoyuzhouFMMediaContent(url string, document *goquery.Document) (string, string, string) {
	audioUrl := ""
	document.Find("meta[property='og:audio']").Each(func(i int, s *goquery.Selection) {
		if content, exists := s.Attr("content"); exists {
			audioUrl = content
		}

	})
	return audioUrl, audioUrl, "audio"
}
