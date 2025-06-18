package templates

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/beclab/article-extractor/readability"
)

func (t *Template) OKjikeScrapContent(document *goquery.Document) string {
	contents := ""
	document.Find("div.info").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})
	document.Find("div.post-wrap").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})

	return contents
}

func (t *Template) OKjikeScrapMetaData(document *goquery.Document) (string, string) {
	author := ""
	published_at := ""

	document.Find("div.title").Each(func(i int, s *goquery.Selection) {
		author = s.Text()
	})
	return author, published_at
}

func (t *Template) OKjikePublishedAtTimeFromScriptMetadata(document *goquery.Document) int64 {

	var publishedAt int64 = 0

	scriptSelector := "script[type=\"application/json\"]"
	document.Find(scriptSelector).Each(func(i int, s *goquery.Selection) {
		scriptContent := strings.TrimSpace(s.Text())
		var metaData map[string]interface{}
		unmarshalErr := json.Unmarshal([]byte(scriptContent), &metaData)
		if unmarshalErr != nil || len(metaData) < 1 {
			log.Printf("convert  unmarshalError %v", unmarshalErr)
		}
		if props, ok := metaData["props"]; ok {
			if pageProps, ok := props.(map[string]interface{})["pageProps"]; ok {
				if post, ok := pageProps.(map[string]interface{})["post"]; ok {
					if createdAt, ok := post.(map[string]interface{})["createdAt"]; ok {
						dateObj, err := readability.ParseTime(createdAt.(string))
						if err == nil {
							publishedAt = dateObj.Unix()
						}
					}
				}
			}
		}

	})

	return publishedAt
}
