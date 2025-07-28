package templates

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func varietyContentExtractor(document *goquery.Document) string {
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

func varietyScrapAuthor(doc *goquery.Document) string {
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

func (t *Template) VarietyExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := varietyContentExtractor(document)
	author := varietyScrapAuthor(document)
	return content, author, 0, "", "", ""
}
