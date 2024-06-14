package templates

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type ScreenRantMetaData struct {
	Context          string `json:"@context"`
	Type             string `json:"@type"`
	MainEntityOfPage struct {
		Type string `json:"@type"`
		ID   string `json:"@id"`
	} `json:"mainEntityOfPage"`
	Headline string `json:"headline"`
	Image    struct {
		Type        string `json:"@type"`
		ContentURL  string `json:"contentUrl"`
		CreditText  string `json:"creditText"`
		Description string `json:"description"`
		Height      string `json:"height"`
		Width       string `json:"width"`
	} `json:"image"`
	DatePublished time.Time `json:"datePublished"`
	DateModified  time.Time `json:"dateModified"`
	Author        []struct {
		Type        string `json:"@type"`
		ID          string `json:"@id"`
		Name        string `json:"name"`
		URL         string `json:"url"`
		Description string `json:"description"`
		JobTitle    string `json:"jobTitle"`
		KnowsAbout  []struct {
			Type string `json:"@type"`
			Name string `json:"name"`
		} `json:"knowsAbout"`
		Image  string   `json:"image"`
		SameAs []string `json:"sameAs"`
	} `json:"author"`
	Publisher struct {
		Type                 string `json:"@type"`
		ID                   string `json:"@id"`
		Name                 string `json:"name"`
		URL                  string `json:"url"`
		Description          string `json:"description"`
		PublishingPrinciples string `json:"publishingPrinciples"`
		FoundingDate         string `json:"foundingDate"`
		SameAs               []any  `json:"sameAs"`
		Logo                 struct {
			Type   string `json:"@type"`
			URL    string `json:"url"`
			Height string `json:"height"`
			Width  string `json:"width"`
		} `json:"logo"`
	} `json:"publisher"`
	ArticleSection      []string `json:"articleSection"`
	Description         string   `json:"description"`
	IsAccessibleForFree string   `json:"isAccessibleForFree"`
	HasPart             []struct {
		Type                string `json:"@type"`
		IsAccessibleForFree string `json:"isAccessibleForFree"`
		CSSSelector         string `json:"cssSelector"`
	} `json:"hasPart"`
}

func (t *Template) ScreenrantScrapContent(document *goquery.Document) string {
	contents := ""

	document.Find("div.adsninja-ad-zone,div.active-content").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})
	document.Find("div.heading_image,section.article-body").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func extractScreenratAuthors(doc *goquery.Document) string {
	var authors []string

	doc.Find("meta[property='article:author']").Each(func(i int, s *goquery.Selection) {
		// 获取content属性的值
		if author, exists := s.Attr("content"); exists {
			authors = append(authors, author)
		}
	})
	var authorsString string = ""
	if len(authors) != 0 {
		authorsString = strings.Join(authors, " & ")

	}
	return authorsString;
}

func (t *Template) ScreenrantScrapMetaData(document *goquery.Document) (string, string) {

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
			var firstTypeMetaData ScreenRantMetaData
			unmarshalErr := json.Unmarshal([]byte(scriptContent), &firstTypeMetaData)
			if unmarshalErr != nil {
				log.Printf("convert ScreenMetadata unmarshalError %v", unmarshalErr)
				return
			}
			for _, currentMetadata := range firstTypeMetaData.Author {
				if len(currentMetadata.Name) != 0 {
					if len(author) != 0 {
						author = author + " & "  + currentMetadata.Name
					}else{
						author = currentMetadata.Name
					}
				}
			}
		})
		if author != "" {
			break
		}
	}
	if len(author) == 0 {
		author = extractScreenratAuthors(document)
	}
	log.Printf("author last: %s", author)
	return author, published_at
}

func stringScreenratToTimestamp(timeStr string) (int64, error) {
	t, err := time.Parse("2006/01/02 15:04:05", timeStr)
	if err != nil {
		return 0, err
	}

	return t.Unix(), nil
}

func extractScreentPublishedTime(doc *goquery.Document) (int64, error) {
	s := doc.Find("meta[property='article:published_time']").First()

	timeStr, exists := s.Attr("content")
	if !exists {
		return 0, fmt.Errorf("article:published_time not found")
	}
	fmt.Printf("published screenrat time %s \n",timeStr)
	t, err := time.Parse(time.RFC3339, timeStr)
	if err == nil {
		return t.Unix(), nil

	}else{
		log.Printf("error parsing time: %v \n", err)
	}

	if err != nil {
		publishedAt,convertErr  := stringScreenratToTimestamp(timeStr)
		if convertErr != nil {
			log.Printf("error fuck ************** %v \n",convertErr)
			return 0, convertErr
		}
		return publishedAt,nil
	}
	return t.Unix(), nil
}

func (t *Template) ScreenrantPublishedAtTimeFromScriptMetadata(document *goquery.Document) int64 {
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
			var firstTypeMetaData ScreenRantMetaData
			unmarshalErr := json.Unmarshal([]byte(scriptContent), &firstTypeMetaData)
			if unmarshalErr != nil {
				log.Printf("convert screenMetaData unmarshalError %v", unmarshalErr)
				return

			}
			
			publishedAt = firstTypeMetaData.DatePublished.Unix()
			//publishedAt = firstTypeMetaData[0].DatePublished.Unix()
		})

	}
	if publishedAt == 0 || publishedAt < 0{
		publishedAtConvert,err :=  extractScreentPublishedTime(document)
		if err == nil {
			publishedAt = publishedAtConvert
		}else{
			log.Printf("extract screenPublishedTime err %v", err)

		}

	}
	return publishedAt
}
