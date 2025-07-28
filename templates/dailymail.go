package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func dailymailScrapContent(document *goquery.Document) string {

	contents := ""
	document.Find("div.related-partners,div.related-carousel,div.moduleFull,div.floatRHS").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})
	document.Find("div").Each(func(i int, s *goquery.Selection) {
		if _, exists := s.Attr("data-podcast-container"); exists {
			RemoveNodes(s)
		}
	})
	document.Find("div[itemprop=articleBody]").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) DailymailExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := dailymailScrapContent(document)

	return content, "", 0, "", "", ""
}
