package templates

import (
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type TimeMetaData []struct {
	Context  string `json:"@context"`
	Type     string `json:"@type"`
	Headline string `json:"headline,omitempty"`
	Image    []struct {
		Type              string `json:"@type"`
		URL               string `json:"url"`
		Width             int    `json:"width"`
		Height            int    `json:"height"`
		AssociatedArticle struct {
			Type string `json:"@type"`
			URL  string `json:"url"`
		} `json:"associatedArticle"`
		Headline    string `json:"headline"`
		Description string `json:"description"`
		CreditText  string `json:"creditText"`
	} `json:"image,omitempty"`
	Author []struct {
		Type string `json:"@type"`
		Name string `json:"name"`
	} `json:"author,omitempty"`
	DatePublished  time.Time `json:"datePublished,omitempty"`
	DateModified   string `json:"dateModified,omitempty"`
	Description    string    `json:"description,omitempty"`
	URL            string    `json:"url,omitempty"`
	ArticleSection string    `json:"articleSection,omitempty"`
	Keywords       []string  `json:"keywords,omitempty"`
	ThumbnailURL   string    `json:"thumbnailUrl,omitempty"`
	Publisher      struct {
		Type string `json:"@type"`
		Name string `json:"name"`
		URL  string `json:"url"`
		Logo struct {
			Type   string `json:"@type"`
			URL    string `json:"url"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"logo"`
		FoundingDate string   `json:"foundingDate"`
		SameAs       []string `json:"sameAs"`
	} `json:"publisher,omitempty"`
	ItemListElement []struct {
		Type     string `json:"@type"`
		Position int    `json:"position"`
		Item     struct {
			ID    string `json:"@id"`
			Name  string `json:"name"`
			Image any    `json:"image"`
		} `json:"item"`
	} `json:"itemListElement,omitempty"`
}

func (t *Template) TimeScrapContent(document *goquery.Document) string {
	contents := ""
	document.Find("p").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		if strings.HasPrefix(text, "Read More:") {
			RemoveNodes(s)
		}

	})

	document.Find("div.featured-media,div#article-body-main").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}


func (t *Template) TimeScrapMetaData(document *goquery.Document) (string, string) {

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
			var firstTypeMetaData TimeMetaData
			unmarshalErr := json.Unmarshal([]byte(scriptContent), &firstTypeMetaData)
			if unmarshalErr != nil {
				log.Printf("convert TimeMetaData unmarshalError %v", unmarshalErr)
				return
			}
			for _, currentMetadata := range firstTypeMetaData {
				if len(author) != 0 {
					break
				}
				if currentMetadata.Type != "NewsArticle" {
					continue
				}
				for _,currentAuthor := range currentMetadata.Author {
					if len(currentAuthor.Name) != 0{
						if len(author) != 0 {
							author = author + " & " + currentAuthor.Name
						}else{
							author = currentAuthor.Name
						}
					}
				}

			}
		})
		if author != "" {
			break
		}
	}
	log.Printf("author last: %s", author)
	return author, published_at
}

func (t *Template) TimePublishedAtTimeFromScriptMetadata(document *goquery.Document) int64 {

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
			var firstTypeMetaData TimeMetaData
			unmarshalErr := json.Unmarshal([]byte(scriptContent), &firstTypeMetaData)
			if unmarshalErr != nil {
				log.Printf("convert SkyNewsScrap unmarshalError %v", unmarshalErr)
				return

			}
			for _, currentMetadata := range firstTypeMetaData {
				if publishedAt != 0 {
					break
				}
				if currentMetadata.Type != "NewsArticle" {
					continue
				}
				publishedAt = currentMetadata.DatePublished.Unix()

			}
			//publishedAt = firstTypeMetaData[0].DatePublished.Unix()
		})

	}
	return publishedAt
}