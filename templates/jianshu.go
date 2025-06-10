package templates

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func (t *Template) JianshuScrapContent(document *goquery.Document) string {
	contents := ""

	document.Find("article").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) JianshuScrapMetaData(document *goquery.Document) (string, string) {

	author := ""
	published_at := ""

	document.Find("script[type='application/json']").Each(func(i int, s *goquery.Selection) {
		scriptContent := strings.TrimSpace(s.Text())
		var metaData map[string]interface{}
		unmarshalErr := json.Unmarshal([]byte(scriptContent), &metaData)
		if unmarshalErr != nil {
			log.Printf("convert  unmarshalError %v", unmarshalErr)
		}
		if props, ok := metaData["props"]; ok {
			if initialState, ok := props.(map[string]interface{})["initialState"]; ok {
				if note, ok := initialState.(map[string]interface{})["note"]; ok {
					if noteData, ok := note.(map[string]interface{})["data"]; ok {
						if user, ok := noteData.(map[string]interface{})["user"]; ok {
							if nickname, ok := user.(map[string]interface{})["nickname"]; ok {
								author = nickname.(string)
							}
						}
					}
				}
			}
		}
	})
	return author, published_at
}
