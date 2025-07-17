package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func filmSchoolRejectsScrapContent(document *goquery.Document) string {

	contents := ""
	document.Find("section.recommended").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})
	document.Find("article.article").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) FilmSchoolRejectsExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := filmSchoolRejectsScrapContent(document)

	return content, "", 0, "", "", ""
}
