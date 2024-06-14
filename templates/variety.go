package templates

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func (t *Template) VarietyScrapContent(document *goquery.Document) string {
	contents := ""
	document.Find("aside,div.article-tags,div.o-comments-link").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})
	document.Find("img").Each(func(i int, img *goquery.Selection) {
		src, exists := img.Attr("data-lazy-src")
		if exists {
			img.SetAttr("src", src)
		}
	})

	document.Find("div.article-header__feature,div.vy-cx-page-content").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func extractVarietyAuthors(doc *goquery.Document) string {
	var authors []string

	doc.Find("meta[name='author']").Each(func(i int, s *goquery.Selection) {
		if author, exists := s.Attr("content"); exists {
			if !checkStrArrContains(authors, author) {
				authors = append(authors, author)
			}
		}
	})
	var authorsString string = ""
	if len(authors) != 0 {
		authorsString = strings.Join(authors, " & ")

	}
	return authorsString
}

func (t *Template) VarietyScrapMetaData(document *goquery.Document) (string, string) {
	published_at := ""
	author := extractVarietyAuthors(document)

	return author, published_at
}

func extractVarietyPublishedTime(doc *goquery.Document) (int64, error) {
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

func (t *Template) VarietyPublishedAtTimeFromScriptMetadata(document *goquery.Document) int64 {
	var publishedAt int64 = 0

	publishedAtConvert, err := extractVarietyPublishedTime(document)
	if err == nil {
		publishedAt = publishedAtConvert
	} else {
		log.Printf("extract screenPublishedTime err %v", err)

	}

	return publishedAt
}
