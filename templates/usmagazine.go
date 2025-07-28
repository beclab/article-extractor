package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func usmagazineContentExtractor(document *goquery.Document) string {
	contents := ""

	document.Find("div#news-block,div.link-related,div.in-this-article").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)

	})
	document.Find("div.article-content").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) UsmagazineExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := usmagazineContentExtractor(document)
	return content, "", 0, "", "", ""
}
