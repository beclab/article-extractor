package templates

import (
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

func v2exContentExtractor(document *goquery.Document) string {
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

func v2exScrapAuthor(document *goquery.Document) string {
	author := ""
	document.Find("meta[name='description']").Each(func(i int, s *goquery.Selection) {
		if content, exists := s.Attr("content"); exists {
			re := regexp.MustCompile(`@(\w+)`)
			matches := re.FindAllString(content, -1)
			if len(matches) > 0 {
				author = matches[0][1:]
			}

		}
	})
	return author
}

func (t *Template) V2exExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := v2exContentExtractor(document)
	author := v2exScrapAuthor(document)
	return content, author, 0, "", "", ""
}
