package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func techCrunchScrapContent(document *goquery.Document) string {
	contents := ""
	document.Find("h1.article__title,div.article__byline").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})
	document.Find("article.article-container").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) TechCrunchExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := techCrunchScrapContent(document)
	return content, "", 0, "", "", ""
}
