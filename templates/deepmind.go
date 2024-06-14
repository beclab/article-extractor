package templates

import (
	"log"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func (t *Template) DeepMindScrapContent(document *goquery.Document) string {

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

func (t *Template) DeepMindScrapMetaData(document *goquery.Document) (string, string) {

	author := ""
	published_at := ""
	document.Find("dd.article-cover__authors").Each(func(index int, item *goquery.Selection) {
		p := item.Find("p")
		author = p.Text()
	})
	return author, published_at
}

func (t *Template) DeepMindPublishedAtTimeFromScriptMetadata(document *goquery.Document) int64 {

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

	// 解析日期格式

	return publishedAt
}
