package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func kotakuScrapContent(document *goquery.Document) string {
	contents := ""

	document.Find("div.js_related-stories-inset,div.js_ad-dynamic,div.instream-native-video,div.js_related-stories-inset-mobile").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})
	document.Find("div.js_post-content").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) KotakuExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := kotakuScrapContent(document)
	return content, "", 0, "", "", ""
}
