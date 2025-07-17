package templates

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func a16ZCrptoScrapAuthor(document *goquery.Document) string {

	author := ""
	var authors []string

	document.Find("meta[name='parsely-author']").Each(func(i int, s *goquery.Selection) {
		if content, exists := s.Attr("content"); exists {
			authors = append(authors, content)
		}
	})
	if len(authors) > 0 {
		author = strings.Join(authors, " & ")
	}

	return author
}

func (t *Template) A16ZCrptoExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	author := a16ZCrptoScrapAuthor(document)
	return "", author, 0, "", "", ""
}
