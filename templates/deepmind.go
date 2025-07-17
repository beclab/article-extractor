package templates

import (
	"log"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func deepMindScrapContent(document *goquery.Document) string {
	contents := ""
	document.Find("div.article-cover__header,aside.related-posts").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})
	document.Find("article").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func deepMindScrapPublishedAt(document *goquery.Document) int64 {
	var publishedAt int64 = 0
	document.Find("dd.article-cover__date").Each(func(index int, item *goquery.Selection) {
		timeTag := item.Find("time")
		datetimeAttr, _ := timeTag.Attr("datetime")
		parsedDate, err := time.Parse("2006-01-02", datetimeAttr)
		if err != nil {
			log.Printf("Error parsing date: %v", err)
			return
		}
		publishedAt = parsedDate.Unix()
	})
	return publishedAt
}

func (t *Template) DeepMindExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := deepMindScrapContent(document)
	author := ""
	document.Find("dd.article-cover__authors").Each(func(index int, item *goquery.Selection) {
		p := item.Find("p")
		author = p.Text()
	})
	publishedAt := deepMindScrapPublishedAt(document)
	return content, author, publishedAt, "", "", ""
}
