package templates

import "github.com/PuerkitoBio/goquery"

func theStarScrapContent(document *goquery.Document) string {
	contents := ""
	document.Find("div.hidden-print").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})
	document.Find("div.asset-photo,div#article-body").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) TheStarExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := theStarScrapContent(document)
	return content, "", 0, "", "", ""
}
