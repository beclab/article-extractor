package templates

import "github.com/PuerkitoBio/goquery"

func telegraphScrapContent(document *goquery.Document) string {
	contents := ""
	document.Find("header,aside").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})
	document.Find("article.grid").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) TelegraphExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := telegraphScrapContent(document)
	return content, "", 0, "", "", ""
}
