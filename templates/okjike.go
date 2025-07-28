package templates

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/beclab/article-extractor/readability"
)

func okjikeScrapContent(document *goquery.Document) string {
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

func okjikeScrapPublishedAt(document *goquery.Document) int64 {
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

func (t *Template) OKjikeExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := okjikeScrapContent(document)
	author := ""
	document.Find("div.title").Each(func(i int, s *goquery.Selection) {
		author = s.Text()
	})
	publishedAt := okjikeScrapPublishedAt(document)
	return content, author, publishedAt, "", "", ""
}
