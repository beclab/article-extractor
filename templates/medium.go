package templates

import "github.com/PuerkitoBio/goquery"

func mediumScrapContent(document *goquery.Document) string {
	document.Find("h1.pw-post-title,div.speechify-ignore").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})

	contents := ""

	document.Find("section").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) MediumExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := mediumScrapContent(document)
	return content, "", 0, "", "", ""
}
