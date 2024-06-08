package processor

import (
	"encoding/json"
	"log"
	"plugin"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/beclab/article-extractor/readability"
)

func getPluginsContent(pluginsPath, entryUrl string, doc *goquery.Document) string {
	log.Printf(`plugins Path %s`, pluginsPath)
	if pluginsPath == "" {
		return ""
	}
	p, err := plugin.Open(pluginsPath)
	if err != nil {
		log.Printf(`open plugins error %s,%v`, pluginsPath, err)
		return ""
	}
	templateExtr, err := p.Lookup("ContentTemplateExtractor")
	if err != nil {
		log.Printf(`lookup plugins error %v`, err)
		return ""
	}
	tempE, ok := templateExtr.(func(string, *goquery.Document) string)
	if !ok {
		log.Printf(`find plugins error %v`, err)
		return ""
	}
	content := tempE(entryUrl, doc)
	return content
}

func getPluginsAuthor(pluginsPath, entryUrl string, doc *goquery.Document) string {
	log.Printf(`plugins Path %s`, pluginsPath)
	if pluginsPath == "" {
		return ""
	}
	p, err := plugin.Open(pluginsPath)
	if err != nil {
		log.Printf(`open plugins error %s,%v`, pluginsPath, err)
		return ""
	}
	templateExtr, err := p.Lookup("AuthorTemplateExtractor")
	if err != nil {
		log.Printf(`lookup plugins error %v`, err)
		return ""
	}
	tempE, ok := templateExtr.(func(string, *goquery.Document) string)
	if !ok {
		log.Printf(`find plugins error %v`, err)
		return ""
	}
	content := tempE(entryUrl, doc)
	return content
}

func getPluginsPublishedAtTemplate(pluginsPath, entryUrl string, doc *goquery.Document) int64 {
	log.Printf(`plugins Path %s`, pluginsPath)
	if pluginsPath == "" {
		return 0
	}
	p, err := plugin.Open(pluginsPath)
	if err != nil {
		log.Printf(`open plugins error %s,%v`, pluginsPath, err)
		return 0
	}
	templateExtr, err := p.Lookup("PublishedAtTemplateExtractor")
	if err != nil {
		log.Printf(`lookup plugins error %v`, err)
		return 0
	}
	tempE, ok := templateExtr.(func(string, *goquery.Document) int64)
	if !ok {
		log.Printf(`find plugins error %v`, err)
		return 0
	}
	content := tempE(entryUrl, doc)
	return content
}

func getPluginsPostContent(pluginsPath, entryUrl, content string) string {
	log.Printf(`plugins Path %s`, pluginsPath)
	if pluginsPath == "" {
		return ""
	}
	p, err := plugin.Open(pluginsPath)
	if err != nil {
		log.Printf(`open plugins error %s,%v`, pluginsPath, err)
		return ""
	}
	templateExtr, err := p.Lookup("PostContentTemplateExtractor")
	if err != nil {
		log.Printf(`lookup plugins error %v`, err)
		return ""
	}
	tempE, ok := templateExtr.(func(string, string) string)
	if !ok {
		log.Printf(`find plugins error %v`, err)
		return ""
	}
	postContent := tempE(entryUrl, content)
	return postContent
}

func checkStrArrContains(arr []string, e string) bool {
	for _, a := range arr {
		if strings.TrimSpace(a) == strings.TrimSpace(e) {
			return true
		}
	}
	return false
}

func ScrapContentUseRules(document *goquery.Document, rules string) (string, error) {
	contents := ""
	document.Find(rules).Each(func(i int, s *goquery.Selection) {
		var content string

		content, _ = goquery.OuterHtml(s)
		contents += content
	})

	return contents, nil
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
	return author
}

func ScrapAutoPublishedAtTimeMetaData(doc *goquery.Document) int64 {
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

func GetPublishedAtTimestampForWechat(rawContent string, url string) int64 {
	var publishedAtTimestamp int64 = 0
	re := regexp.MustCompile(`var oriCreateTime = '(\d+)';`)
	match := re.FindStringSubmatch(rawContent)
	if len(match) > 1 {
		timestamp, err := strconv.ParseInt(match[1], 10, 64)
		if err != nil {
			log.Printf("can not parse timestamp [%s] for entry [%s]", match[1], url)
			return publishedAtTimestamp
		}
		publishedAtTimestamp = timestamp
	} else {
		log.Printf("can not find timestamp for entry [%s]", url)
		return publishedAtTimestamp
	}
	return publishedAtTimestamp

}
