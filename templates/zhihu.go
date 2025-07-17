package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func (t *Template) ZhihuExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	contents := ""
	document.Find("div.Post-RichTextContainer").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})

	author := ""
	document.Find("meta[itemprop='name']").Each(func(i int, s *goquery.Selection) {
		if content, exists := s.Attr("content"); exists {
			author = content
		}
	})
	return contents, author, 0, "", "", ""
}
