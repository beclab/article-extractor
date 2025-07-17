package templates

import (
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func cbsSportsScrapContent(document *goquery.Document) string {
	contents := ""
	document.Find("div.Article-body").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func cbsSportsScrapAuthor(document *goquery.Document) string {

	author := ""
	var authors []string
	document.Find("a.ArticleAuthor-name--link").Each(func(index int, item *goquery.Selection) {
		authorName := item.Text()
		authors = append(authors, authorName)
	})
	author = strings.Join(authors, " & ")
	if len(author) == 0 {
		document.Find("span.ArticleAuthor-nameText").Each(func(index int, item *goquery.Selection) {
			authorName := strings.TrimSpace(item.Text())
			authors = append(authors, authorName)
		})
		authorsString := strings.Join(authors, " & ")
		author = authorsString
	}

	return author
}

func cbsSportScrapPublishedAt(document *goquery.Document) int64 {
	var publishedAt int64 = 0
	document.Find("time.TimeStamp").Each(func(index int, item *goquery.Selection) {
		if publishedAt != 0 {
			return
		}
		datetime, exists := item.Attr("datetime")
		if exists {
			t, err := time.Parse("2006-01-02 15:04:05 MST", datetime)
			if err != nil {
				return
			}
			publishedAt = t.Unix()

		}
	})

	return publishedAt
}

func (t *Template) CBSSportExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := cbsSportsScrapContent(document)
	author := cbsSportsScrapAuthor(document)
	publishedAt := cbsSportScrapPublishedAt(document)
	return content, author, publishedAt, "", "", ""
}
