package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func (t *Template) YcombinatorExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	contents := ""
	document.Find("div.prose").Each(func(i int, s *goquery.Selection) {
		var content string
		prev := s.Prev()
		if prev.Length() > 0 {
			content, _ = goquery.OuterHtml(prev)
			contents += content
		}

		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents, "", 0, "", "", ""
}
