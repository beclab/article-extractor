package templates

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func newatlasScrapContent(document *goquery.Document) string {
	contents := ""

	document.Find("div.FullscreenCarousel-cover,div.ArticlePage-articleContainer").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

type NewatlasMetaData struct {
	Context string `json:"@context"`
	Graph   []struct {
		Type string `json:"@type"`
		Name string `json:"name"`
	} `json:"@graph"`
}

func newatlasScrapAuthor(doc *goquery.Document) string {
	author := ""
	scriptSelector := "script[type=\"application/ld+json\"]"
	doc.Find(scriptSelector).Each(func(i int, s *goquery.Selection) {
		scriptContent := strings.TrimSpace(s.Text())
		var metaData NewatlasMetaData
		unmarshalErr := json.Unmarshal([]byte(scriptContent), &metaData)
		if unmarshalErr != nil {
			log.Printf("convert SkyNewsScrap unmarshalError %v", unmarshalErr)
		} else {
			for _, graphData := range metaData.Graph {
				if graphData.Type == "Person" {
					author = graphData.Name
					return
				}
			}
		}
	})
	return author
}

func (t *Template) NewatlasExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := newatlasScrapContent(document)
	author := newatlasScrapAuthor(document)
	return content, author, 0, "", "", ""
}
