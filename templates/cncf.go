package templates

import (
	"encoding/json"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type CNCFMetaData struct {
	Context string `json:"@context"`
	Graph   []struct {
		Type            string `json:"@type"`
		ID              string `json:"@id"`
		URL             string `json:"url"`
		Name            string `json:"name"`
		InLanguage      string `json:"inLanguage"`
		Publisher struct {
			Type   string   `json:"@type"`
			ID     string   `json:"@id"`
			Name   string   `json:"name"`
			URL    string   `json:"url"`
			SameAs []string `json:"sameAs"`
			Logo   struct {
				Type        string `json:"@type"`
				URL         string `json:"url"`
				ContentURL  string `json:"contentUrl"`
				Width       int    `json:"width"`
				Height      int    `json:"height"`
				ContentSize string `json:"contentSize"`
			} `json:"logo"`
		} `json:"publisher,omitempty"`
		Description string `json:"description,omitempty"`
		IsPartOf    struct {
			ID string `json:"@id"`
		} `json:"isPartOf,omitempty"`
		Breadcrumb struct {
			Type            string `json:"@type"`
			ID              string `json:"@id"`
			ItemListElement []struct {
				Type     string `json:"@type"`
				Position int    `json:"position"`
				Item     string `json:"item,omitempty"`
				Name     string `json:"name"`
			} `json:"itemListElement"`
		} `json:"breadcrumb,omitempty"`
		DatePublished string `json:"datePublished,omitempty"`
		DateModified  string `json:"dateModified,omitempty"`
		Author        struct {
			Type string `json:"@type"`
			ID   string `json:"@id"`
			Name string `json:"name"`
		} `json:"author,omitempty"`
	} `json:"@graph"`
}

func (t *Template) CNCFScrapContent(document *goquery.Document) string {

	contents := ""
	document.Find("div.social-share,figure.wp-block-embed-twitter,div.post-author").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})

	document.Find("article.container").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) CNCFScrapMetaData(document *goquery.Document) (string, string) {

	author := ""
	published_at := ""
	/*
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
			var firstTypeMetaData CNCFMetaData
			unmarshalErr := json.Unmarshal([]byte(scriptContent), &firstTypeMetaData)
			if unmarshalErr != nil {
				log.Printf("convert SkyNewsScrap unmarshalError %v", unmarshalErr)
				return
			}
			for _, currentGraph := range firstTypeMetaData.Graph {
				if len(currentGraph.Author.Name) != 0 {
					author = author + " & " + currentGraph.Author.Name
				} else {
					author = currentGraph.Author.Name
				}
			}
		})
		if author != "" {
			break
		}
	}*/

	document.Find("span.post-author__author").Each(func(i int, s *goquery.Selection) {
        spanContent := s.Text()
		author = strings.TrimPrefix(spanContent, "By ")

    })

	if len(author) == 0 {
		document.Find("div.post-author").Next().Each(func(i int, s *goquery.Selection) {
			emContent := s.Find("em").Text()
			cleanContent := strings.Replace(emContent, "\u00a0", "", -1) 
			cleanContent = strings.TrimPrefix(cleanContent, "Community post by ")
			cleanContent = strings.TrimPrefix(cleanContent, "Member post by ")
			cleanContent = strings.TrimPrefix(cleanContent,"Community post originally published on Medium by ")
			cleanContent = strings.TrimPrefix(cleanContent,"Project post by ")
			cleanContent = strings.TrimPrefix(cleanContent,"Member post originally published on Greptime’s blog by ")
			extractAuthor,extractSuccess:=extractAuthorFromCNCFName(cleanContent)
			if extractSuccess {
				author = extractAuthor
			}else{
				author = cleanContent
			}
		})

	}

	return author, published_at
}

func parseCNCFDateTimeToTimestamp(dateTimeStr string) (int64, error) {
	dateTimeStr =  addSecondsIfMissing(dateTimeStr)
	log.Printf("-------------- %s",dateTimeStr)
	layout := "2006-01-02T15:04:05-07:00" 
	t, err := time.Parse(layout, dateTimeStr)
	if err != nil {
		log.Printf("parse cncf time err %v",err)
		return 0,err
	}
	return t.Unix(), nil
}


func addSecondsIfMissing(dateTimeStr string) string {
    re := regexp.MustCompile(`^(\d{4}-\d{2}-\d{2}T\d{2}:\d{2})\+(\d{2}:\d{2})$`)

    if re.MatchString(dateTimeStr) {
        dateTimeStr = re.ReplaceAllString(dateTimeStr, "${1}:00+${2}")
    }
    return dateTimeStr
}
func extractAuthorFromCNCFName(input string) (string, bool) {
	re := regexp.MustCompile(`Member post originally published on ([^’]+)’s blog`)
	matches := re.FindStringSubmatch(input)
	if len(matches) >= 2 {
		return matches[1], true
	}
	return "", false
}

func (t *Template) CNCFPublishedAtTimeFromScriptMetadata(document *goquery.Document) int64 {

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
			var firstTypeMetaData CNCFMetaData
			unmarshalErr := json.Unmarshal([]byte(scriptContent), &firstTypeMetaData)
			if unmarshalErr != nil {
				log.Printf("convert cncf unmarshalError %v", unmarshalErr)
				return

			}

			for _, currentGraph := range firstTypeMetaData.Graph {
				if publishedAt != 0 {
					return
				}
				parseTime,parseTimeErr:=parseCNCFDateTimeToTimestamp(currentGraph.DatePublished)
				if parseTimeErr != nil {
					log.Printf("parset CNCF time to timetamp error %v",parseTimeErr)
					continue
				}
				publishedAt =parseTime
			}
			//publishedAt = firstTypeMetaData[0].DatePublished.Unix()
		})

	}
	return publishedAt
}
