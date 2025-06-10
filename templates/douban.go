package templates

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/beclab/article-extractor/readability"
)

func (t *Template) DoubanScrapContent(document *goquery.Document) string {

	contents := ""
	document.Find("div.topic-richtext,div.review-content").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) DoubanScrapMetaData(document *goquery.Document) (string, string) {

	author := ""
	published_at := ""

	document.Find("script[type='application/ld+json']").Each(func(i int, s *goquery.Selection) {
		scriptContent := strings.TrimSpace(s.Text())
		var metaData map[string]interface{}
		unmarshalErr := json.Unmarshal([]byte(scriptContent), &metaData)
		if unmarshalErr != nil {
			log.Printf("convert  unmarshalError %v", unmarshalErr)
		}
		if authors, ok := metaData["author"]; ok {
			if name, ok := authors.(map[string]interface{})["name"]; ok {
				author = name.(string)
			}
		}
	})
	if author == "" {
		document.Find("span.from").Each(func(i int, s *goquery.Selection) {
			author = s.Text()
		})

	}
	return author, published_at
}

func (t *Template) DoubanPublishedAtTimeFromScriptMetadata(document *goquery.Document) int64 {

	var publishedAt int64 = 0

	document.Find("span.create-time").Each(func(i int, s *goquery.Selection) {
		publishTimes := strings.TrimSpace(s.Text())
		dateObj, err := readability.ParseTime(publishTimes)
		if err == nil {
			publishedAt = dateObj.Unix()
		}
	})

	if publishedAt == 0 {
		scriptSelector := "script[type=\"application/ld+json\"]"
		document.Find(scriptSelector).Each(func(i int, s *goquery.Selection) {
			scriptContent := strings.TrimSpace(s.Text())
			var metaData map[string]interface{}
			unmarshalErr := json.Unmarshal([]byte(scriptContent), &metaData)
			if unmarshalErr != nil || len(metaData) < 1 {
				log.Printf("convert  unmarshalError %v", unmarshalErr)
			}
			if publishTimes, ok := metaData["datePublished"]; ok {
				dateObj, err := readability.ParseTime(publishTimes.(string))
				if err == nil {
					publishedAt = dateObj.Unix()
				}
			}

		})
	}
	return publishedAt
}
