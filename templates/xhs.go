package templates

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func (t *Template) XhsScrapMetaData(document *goquery.Document) (string, string) {
	author := ""
	published_at := ""

	var jsonData string
	document.Find("script").Each(func(i int, s *goquery.Selection) {
		scriptContent := s.Text()
		if strings.Contains(scriptContent, "window.__INITIAL_STATE__") {
			parts := strings.SplitN(scriptContent, "=", 2)
			if len(parts) == 2 {
				jsonData = strings.TrimSpace(parts[1])
				return
			}
		}
	})
	jsonData = strings.ReplaceAll(jsonData, "undefined", "null")
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(jsonData), &result); err != nil {
		log.Printf("json unmarshal %v", err)
	}
	if note, ok := result["note"]; ok {
		if noteDetailMap, ok := note.(map[string]interface{})["noteDetailMap"]; ok {
			for key, item := range noteDetailMap.(map[string]interface{}) {
				if itemNote, ok := item.(map[string]interface{})["note"]; ok {
					if user, ok := itemNote.(map[string]interface{})["user"]; ok {
						if nickname, ok := user.(map[string]interface{})["nickname"]; ok {
							author = nickname.(string)
						}
					}
				}
				fmt.Printf("xhs Key: %s\n", key)
			}
		}
	}
	return author, published_at
}

func (t *Template) XhsPublishedAtTimeFromScriptMetadata(document *goquery.Document) int64 {
	var publishedAt int64 = 0
	var jsonData string
	document.Find("script").Each(func(i int, s *goquery.Selection) {
		scriptContent := s.Text()
		if strings.Contains(scriptContent, "window.__INITIAL_STATE__") {
			parts := strings.SplitN(scriptContent, "=", 2)
			if len(parts) == 2 {
				jsonData = strings.TrimSpace(parts[1])
				return
			}
		}
	})
	jsonData = strings.ReplaceAll(jsonData, "undefined", "null")
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(jsonData), &result); err != nil {
		log.Printf("json unmarshal %v", err)
	}
	if note, ok := result["note"]; ok {
		if noteDetailMap, ok := note.(map[string]interface{})["noteDetailMap"]; ok {
			for key, item := range noteDetailMap.(map[string]interface{}) {
				if itemNote, ok := item.(map[string]interface{})["note"]; ok {
					if lastUpdateTime, ok := itemNote.(map[string]interface{})["lastUpdateTime"]; ok {
						publishedAt = int64(lastUpdateTime.(float64) / 1000)
					}
				}
				fmt.Printf("xhs Key: %s\n", key)
			}
		}
	}
	return publishedAt
}

func (t *Template) XhsScrapContent(document *goquery.Document) string {
	contents := ""

	document.Find("meta[name='og:image']").Each(func(i int, s *goquery.Selection) {
		if content, exists := s.Attr("content"); exists {
			contents = contents + "<img src='" + content + "' /> <br>"
		}
	})
	document.Find("span.note-text").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}
