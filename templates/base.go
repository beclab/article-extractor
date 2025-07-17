package templates

import (
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/beclab/article-extractor/readability"
)

type Template struct{}

type ExtractorFileInfo struct {
	DownloadURL string
	FileName    string
	FileType    string
}

func GetArticleByDivClass(document *goquery.Document) string {
	content := ""
	document.Find("div.entry-content,div.content-entry,div.article-detail,div.entry,div.entry__content,div.article__content,div.articleContent").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		len := usefulContentLen(text)
		if len > 300 {
			s.Children().Each(func(i int, childsection *goquery.Selection) {
				checkUnusedlDiv(childsection)
			})
			content, _ = goquery.OuterHtml(s)
		}
	})

	return content
}

func ScrapAuthorMetaData(doc *goquery.Document) string {
	author := ""
	scriptSelector := "script[type=\"application/ld+json\"]"
	doc.Find(scriptSelector).Each(func(i int, s *goquery.Selection) {
		var authors []string
		scriptContent := strings.TrimSpace(s.Text())
		var metaData map[string]interface{}
		unmarshalErr := json.Unmarshal([]byte(scriptContent), &metaData)
		if unmarshalErr != nil {
			log.Printf("convert  unmarshalError %v", unmarshalErr)
		}
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

	})

	if author == "" {
		var authors []string
		doc.Find("meta[name='author'],meta[property='author']").Each(func(i int, s *goquery.Selection) {
			if author, exists := s.Attr("content"); exists {
				if !checkStrArrContains(authors, author) {
					authors = append(authors, author)
				}
			}
		})
		if len(authors) != 0 {
			author = strings.Join(authors, " & ")
		}
	}
	if author == "" {
		//https://forum.olares.cn/t/topic/73
		doc.Find("span[itemprop='author']").Each(func(i int, s *goquery.Selection) {
			author = strings.TrimSpace(s.Text())
		})
	}

	return author
}

func ScrapPublishedAtTimeMetaData(doc *goquery.Document) int64 {
	s := doc.Find("meta[property='article:published_time']").First()
	timeStr, exists := s.Attr("content")
	if exists {
		ptime, parseErr := readability.ParseTime(timeStr)
		if parseErr == nil {
			return ptime.Unix()
		}
	}
	var publishedAtTimestamp int64
	scriptSelector := "script[type=\"application/ld+json\"]"
	doc.Find(scriptSelector).Each(func(i int, s *goquery.Selection) {
		scriptContent := strings.TrimSpace(s.Text())
		var metaData map[string]interface{}
		unmarshalErr := json.Unmarshal([]byte(scriptContent), &metaData)
		if unmarshalErr != nil {
			log.Printf("convert  unmarshalError %v", unmarshalErr)
		}
		if publishedData, ok := metaData["datePublished"]; ok {
			currentPublishedAtStr := publishedData.(string)
			ptime, parseErr := readability.ParseTime(currentPublishedAtStr)
			if parseErr == nil {
				publishedAtTimestamp = ptime.Unix()
				return
			}
		}

	})
	return publishedAtTimestamp
}
func checkUnusedlDiv(s *goquery.Selection) bool {
	is_remove := false
	node := s.Get(0)
	d1 := node.Data
	content := s.Text()
	if strings.ToLower(strings.TrimSpace(content)) == "advertisement" {
		RemoveNodes(s)
		is_remove = true
	}
	if d1 == "ul" && len(content) == 0 {
		RemoveNodes(s)
		is_remove = true
	}
	return is_remove
}

func checkStrArrContains(arr []string, e string) bool {
	for _, a := range arr {
		if strings.TrimSpace(a) == strings.TrimSpace(e) {
			return true
		}
	}
	return false
}

func usefulContentLen(text string) int {
	content := strings.Replace(text, " ", "", -1)
	content = strings.Replace(content, "\n", "", -1)
	content = strings.Replace(content, "\t", "", -1)
	return len(content)
}

func RemoveNodes(s *goquery.Selection) {
	s.Each(func(i int, s *goquery.Selection) {
		parent := s.Parent()
		if parent.Length() > 0 {
			parent.Get(0).RemoveChild(s.Get(0))
		}
	})
}

func StringToTimestamp(timeStr string) (int64, error) {
	t, err := time.Parse("2006/01/02 15:04:05", timeStr)
	if err != nil {
		return 0, err
	}

	return t.Unix(), nil
}
