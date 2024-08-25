package templates

import (
	"time"

	"github.com/PuerkitoBio/goquery"
)

func (t *Template) ACFunScrapContent(document *goquery.Document) string {
	contents := ""

	document.Find("div.description-container").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) ACFunScrapMetaData(document *goquery.Document) (string, string) {

	author := ""
	published_at := ""
	document.Find("div.up-info>a").Each(func(i int, s *goquery.Selection) {
		author = s.Text()
	})

	return author, published_at
}

func (t *Template) ACFunMediaContent(url string, document *goquery.Document) (string, string, string) {
	return "", url, "video"
}

func (t *Template) ACFunPublishedAtTimeFromScriptMetadata(document *goquery.Document) int64 {

	var publishedAt int64 = 0

	document.Find("div.publish-time").Each(func(i int, s *goquery.Selection) {
		//发布于 2024-8-21
		publishTimes := s.Text()
		dateObj, err := time.Parse("发布于\u00a02006-1-2", publishTimes)
		if err == nil {
			publishedAt = dateObj.Unix()
		}
	})

	return publishedAt
}
