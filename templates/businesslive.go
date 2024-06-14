package templates

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type BusinessliveMetadata struct {
	Context          string    `json:"@context"`
	Type             string    `json:"@type"`
	MainEntityOfPage string    `json:"mainEntityOfPage"`
	Headline         string    `json:"headline"`
	Description      string    `json:"description"`
	DatePublished    string `json:"datePublished"`
	DateModified     string    `json:"dateModified"`
	Author           struct {
		Type string `json:"@type"`
		Name string `json:"name"`
	} `json:"author"`
	Image []struct {
		Type   string `json:"@type"`
		URL    string `json:"url"`
		Author string `json:"author"`
		Height string `json:"height"`
		Width  string `json:"width"`
	} `json:"image"`
	Publisher struct {
		Type string `json:"@type"`
		Name string `json:"name"`
		URL  string `json:"url"`
		Logo struct {
			Type string `json:"@type"`
			URL  string `json:"url"`
		} `json:"logo"`
		SameAs []string `json:"sameAs"`
	} `json:"publisher"`
}

func (t *Template) BusinessliveScrapMetaData(document *goquery.Document) (string, string) {

	author := ""
	published_at := ""
	scriptSelectorFirst := "head > script[type=\"application/ld+json\"]"
	scriptSelectorSecond := "body > script[type=\"application/ld+json\"]"
	scriptSelectorThird := "script[type=\"application/ld+json\"]"

	scriptSelectorList := make([]string, 100)
	scriptSelectorList = append(scriptSelectorList, scriptSelectorFirst)
	scriptSelectorList = append(scriptSelectorList, scriptSelectorSecond)
	scriptSelectorList = append(scriptSelectorList, scriptSelectorThird)

	for _, scriptSelector := range scriptSelectorList {

		document.Find(scriptSelector).Each(func(i int, s *goquery.Selection) {
			if author != "" {
				return
			}
			scriptContent := strings.TrimSpace(s.Text())
			var firstTypeMetaData BusinessliveMetadata
			unmarshalErr := json.Unmarshal([]byte(scriptContent), &firstTypeMetaData)
			if unmarshalErr != nil {
				log.Printf("convert businesslive unmarshalError %v", unmarshalErr)
				return
			}
			author = firstTypeMetaData.Author.Name
		})
		if author != "" {
			break
		}
	}
	log.Printf("author last: %s", author)
	return author, published_at
}

func (t *Template) BusinesslivePublishedAtTimeFromScriptMetadata(document *goquery.Document) int64 {

	var publishedAt int64 = 0

	scriptSelectorFirst := "head > script[type=\"application/ld+json\"]"
	scriptSelectorSecond := "body > script[type=\"application/ld+json\"]"
	scriptSelectorThird := "script[type=\"application/ld+json\"]"

	scriptSelectorList := make([]string, 100)
	scriptSelectorList = append(scriptSelectorList, scriptSelectorFirst)
	scriptSelectorList = append(scriptSelectorList, scriptSelectorSecond)
	scriptSelectorList = append(scriptSelectorList, scriptSelectorThird)

	for _, scriptSelector := range scriptSelectorList {

		document.Find(scriptSelector).Each(func(i int, s *goquery.Selection) {
			if publishedAt != 0 {
				return
			}
			scriptContent := strings.TrimSpace(s.Text())
			var firstTypeMetaData BusinessliveMetadata
			unmarshalErr := json.Unmarshal([]byte(scriptContent), &firstTypeMetaData)
			if unmarshalErr != nil {
				log.Printf("convert SkyNewsScrap unmarshalError %v", unmarshalErr)
				return

			}
			fmt.Println(firstTypeMetaData.DatePublished)
			parsedTimeStamp,parseTimeStampErr := parseBusinessliveStrTimeToTimestamp(firstTypeMetaData.DatePublished)
			if parseTimeStampErr != nil {
				log.Printf("parse businesslive str time to timestamp err")
				return
			}
			publishedAt = parsedTimeStamp
			// publishedAt = firstTypeMetaData.DatePublished
		})

	}
	return publishedAt
}

func (t *Template) BusinessLiveScrapContent(document *goquery.Document) string {

	contents := ""
	document.Find("div.article-widget-related_articles").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})

	document.Find("div.article-widgets").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents

}

func parseBusinessliveStrTimeToTimestamp(timeStr string) (int64, error) {
	const layout = "2006-01-02T15:04:05-0700"
	t, err := time.Parse(layout, timeStr)
	if err != nil {
		return 0, err // 返回错误和0时间戳
	}
	return t.Unix(), nil
}