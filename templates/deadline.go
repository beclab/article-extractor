package templates

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func (t *Template) DeadlineScrapContent(document *goquery.Document) string {

	contents := ""
	document.Find("div.injected-related-story,section.toaster,div.article-tags,div#comments-loading,h2#comments-title,p.subscribe-to").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})
	document.Find("article").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

type DeadlineMetaData struct {
	Context          string `json:"@context"`
	Type             string `json:"@type"`
	URL              string `json:"url"`
	Name             string `json:"name"`
	MainEntityOfPage struct {
		Type string `json:"@type"`
		ID   string `json:"@id"`
	} `json:"mainEntityOfPage"`
	Headline      string    `json:"headline"`
	DatePublished time.Time `json:"datePublished"`
	DateModified  time.Time `json:"dateModified"`
	ArticleBody   string    `json:"articleBody"`
	Author        struct {
		Type string `json:"@type"`
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"author"`
	Publisher struct {
		Type string `json:"@type"`
		Name string `json:"name"`
		URL  string `json:"url"`
		/*Logo struct {
			Type string `json:"@type"`
			URL  string `json:"url"`
		} `json:"logo"`*/
	} `json:"publisher"`
	IsAccessibleForFree string `json:"isAccessibleForFree"`
	Image               struct {
		Type string `json:"@type"`
		Url  string `json:"url"`
	} `json:"image"`
}

func extractDeadlineAuthors(doc *goquery.Document) string {
	var authors []string

	doc.Find("meta[name='author']").Each(func(i int, s *goquery.Selection) {
		if author, exists := s.Attr("content"); exists {
			authors = append(authors, author)
		}
	})
	var authorsString string = ""
	if len(authors) != 0 {
		authorsString = strings.Join(authors, " & ")

	}
	return authorsString
}

func (t *Template) DeadlineScrapMetaData(document *goquery.Document) (string, string) {
	author := ""
	published_at := ""

	/*scriptSelector := "script[type=\"application/ld+json\"]"
	document.Find(scriptSelector).Each(func(i int, s *goquery.Selection) {
		scriptContent := strings.TrimSpace(s.Text())
		var metaData DeadlineMetaData
		unmarshalErr := json.Unmarshal([]byte(scriptContent), &metaData)
		if unmarshalErr != nil {
			log.Printf("convert ScreenMetadata unmarshalError %v", unmarshalErr)
		}
		author = metaData.Author.Name
		if author != "" {
			return
		}

	})

	if len(author) == 0 {
		author = extractDeadlineAuthors(document)
	}*/
	author = extractDeadlineAuthors(document)
	return author, published_at
}

func extractDeadlinePublishedTime(doc *goquery.Document) (int64, error) {
	s := doc.Find("meta[property='article:published_time']").First()

	timeStr, exists := s.Attr("content")
	if !exists {
		return 0, fmt.Errorf("article:published_time not found")
	}
	t, err := time.Parse(time.RFC3339, timeStr)
	if err == nil {
		return t.Unix(), nil

	} else {
		log.Printf("error parsing time: %v \n", err)
	}

	if err != nil {
		publishedAt, convertErr := StringToTimestamp(timeStr)
		if convertErr != nil {
			return 0, convertErr
		}
		return publishedAt, nil
	}
	return t.Unix(), nil
}

func (t *Template) DeadlinePublishedAtTimeFromScriptMetadata(document *goquery.Document) int64 {
	var publishedAt int64 = 0

	/*scriptSelector := "script[type=\"application/ld+json\"]"

	document.Find(scriptSelector).Each(func(i int, s *goquery.Selection) {
		if publishedAt != 0 {
			return
		}
		scriptContent := strings.TrimSpace(s.Text())
		var metaData DeadlineMetaData
		unmarshalErr := json.Unmarshal([]byte(scriptContent), &metaData)
		if unmarshalErr != nil {
			log.Printf("convert screenMetaData unmarshalError %v", unmarshalErr)

		}

		publishedAt = metaData.DatePublished.Unix()
	})*/

	if publishedAt == 0 || publishedAt < 0 {
		publishedAtConvert, err := extractDeadlinePublishedTime(document)
		if err == nil {
			publishedAt = publishedAtConvert
		} else {
			log.Printf("extract screenPublishedTime err %v", err)

		}

	}
	return publishedAt
}
