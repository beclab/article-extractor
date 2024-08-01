package templates

import (

	//"log"

	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/beclab/article-extractor/readability"
)

type ApNewsMetaData []struct {
	Context       string    `json:"@context"`
	Type          string    `json:"@type"`
	URL           string    `json:"url"`
	DateModified  time.Time `json:"dateModified"`
	DatePublished time.Time `json:"datePublished"`
	Description   string    `json:"description"`
	Image         []struct {
		Context      string `json:"@context"`
		Type         string `json:"@type"`
		Height       int    `json:"height"`
		ThumbnailURL string `json:"thumbnailUrl"`
		URL          string `json:"url"`
		Width        int    `json:"width"`
	} `json:"image"`
	MainEntityOfPage struct {
		Type string `json:"@type"`
		ID   string `json:"@id"`
	} `json:"mainEntityOfPage"`
	Author []struct {
		Context string `json:"@context"`
		Type    string `json:"@type"`
		Name    string `json:"name"`
	} `json:"author"`
	Publisher struct {
		Type string `json:"@type"`
		Name string `json:"name"`
		Logo struct {
			Type string `json:"@type"`
			URL  string `json:"url"`
		} `json:"logo"`
	} `json:"publisher"`
	ArticleSection []string `json:"articleSection"`
	Keywords       []string `json:"keywords"`
	ThumbnailURL   string   `json:"thumbnailUrl"`
	Name           string   `json:"name"`
	Headline       string   `json:"headline"`
}

func (t *Template) ApNewsScrapContent(document *goquery.Document) string {
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

func (t *Template) ApNewsCommonGetPublishedAtTimestamp(document *goquery.Document) int64 {
	var publishedAt int64 = 0
	s := document.Find("meta[property='article:published_time']").First()
	timeStr, exists := s.Attr("content")
	if exists {
		ptime, parseErr := readability.ParseTime(timeStr)
		if parseErr == nil {
			return ptime.Unix()
		}
	}
	return publishedAt

}

func (t *Template) ApnNewsScrapMetaData(document *goquery.Document) (string, string) {
	author := ""
	published_at := ""

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

	return author, published_at
}
