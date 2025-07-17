package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func entrepreneurScrapContent(document *goquery.Document) string {

	contents := ""
	document.Find("article>figure,div.prose").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) EntrepreneurExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := entrepreneurScrapContent(document)

	return content, "", 0, "", "", ""
}
