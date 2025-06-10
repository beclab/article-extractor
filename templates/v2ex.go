package templates

import (
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

func (t *Template) V2exScrapContent(document *goquery.Document) string {
	contents := ""

	document.Find("div.topic_content").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	if contents == "" {
		contents = "<span>.</span>" //no content
	}

	return contents
}

func (t *Template) V2exScrapMetaData(document *goquery.Document) (string, string) {

	author := ""
	published_at := ""

	document.Find("meta[name='description']").Each(func(i int, s *goquery.Selection) {
		if content, exists := s.Attr("content"); exists {
			re := regexp.MustCompile(`@(\w+)`)
			matches := re.FindAllString(content, -1)
			if len(matches) > 0 {
				author = matches[0][1:]
			}

		}
	})
	return author, published_at
}
