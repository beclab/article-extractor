package templates

import (

	//"log"

	"encoding/json"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func apNewsScrapContent(document *goquery.Document) string {
	contents := ""
	document.Find("div.Advertisement,div.Enhancement,div.ActionBar").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})

	document.Find("div.Page-lead>figure,div.RichTextStoryBody").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func apnNewsScrapAuthor(document *goquery.Document) string {
	author := ""
	scriptSelector := "script[type=\"application/ld+json\"]"
	document.Find(scriptSelector).Each(func(i int, s *goquery.Selection) {
		var authors []string
		scriptContent := strings.TrimSpace(s.Text())
		var metaDatas []map[string]interface{}
		unmarshalErr := json.Unmarshal([]byte(scriptContent), &metaDatas)
		if unmarshalErr != nil {
			log.Printf("convert  unmarshalError %v", unmarshalErr)
		}
		for _, metaData := range metaDatas {
			if authorData, ok := metaData["author"]; ok {
				switch authorData.(type) {
				case []interface{}:
					for _, authorDetail := range authorData.([]interface{}) {
						authorMap := authorDetail.(map[string]interface{})
						if authorName, ok := authorMap["name"]; ok {
							a := authorName.(string)
							if !checkStrArrContains(authors, a) {
								authors = append(authors, a)
							}
						}
					}
					if len(authors) > 0 {
						author = strings.Join(authors, " & ")
					}
				case map[string]interface{}:
					authorDetail := authorData.(map[string]interface{})
					if authorName, ok := authorDetail["name"]; ok {
						author = authorName.(string)
					}
				}
			}
		}
		if author != "" {
			return
		}
	})
	if author == "" {
		document.Find("meta[name='gtm-dataLayer']").Each(func(i int, s *goquery.Selection) {
			gtmContent, exists := s.Attr("content")
			if exists {
				var metaData map[string]interface{}
				unmarshalErr := json.Unmarshal([]byte(gtmContent), &metaData)
				if unmarshalErr != nil {
					log.Printf("convert  unmarshalError %v", unmarshalErr)
				}
				if authorData, ok := metaData["author"]; ok {
					author = authorData.(string)
				}
			}
		})
	}

	return author
}

func (t *Template) ApNewsExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := apNewsScrapContent(document)
	author := apnNewsScrapAuthor(document)
	return content, author, 0, "", "", ""
}
