package templates

import "github.com/PuerkitoBio/goquery"

func hunterWalkScrapContent(document *goquery.Document) string {
	contents := ""
	document.Find("div.sharedaddy").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})

	document.Find("div.entry-content").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) HunterWalkExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := hunterWalkScrapContent(document)
	return content, "", 0, "", "", ""
}
