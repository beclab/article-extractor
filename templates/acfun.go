package templates

import (
	"time"

	"github.com/PuerkitoBio/goquery"
)

func acFunScrapContent(document *goquery.Document) string {
	contents := ""

	document.Find("div.description-container").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func acFunScrapPublishedAt(document *goquery.Document) int64 {
	var publishedAt int64 = 0
	document.Find("div.publish-time").Each(func(i int, s *goquery.Selection) {
		publishTimes := s.Text()
		dateObj, err := time.Parse("发布于\u00a02006-1-2", publishTimes)
		if err == nil {
			publishedAt = dateObj.Unix()
		}
	})
	return publishedAt
}

func (t *Template) ACFunExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := acFunScrapContent(document)
	author := ""
	document.Find("div.up-info>a").Each(func(i int, s *goquery.Selection) {
		author = s.Text()
	})
	publishedAt := acFunScrapPublishedAt(document)
	return content, author, publishedAt, "", url, VideoFileType
}
