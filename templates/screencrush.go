package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func screencrushScrapContent(document *goquery.Document) string {
	contents := ""

	document.Find("div.bands-in-town-container,div.tsm-ad,div.branded-app-shortcode-inarticle,iframe").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})
	document.Find("div.content").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) ScreencrushExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := screencrushScrapContent(document)
	return content, "", 0, "", "", ""
}
