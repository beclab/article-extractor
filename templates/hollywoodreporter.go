package templates

import "github.com/PuerkitoBio/goquery"

func hollywoodreporterScrapContent(document *goquery.Document) string {
	contents := ""
	document.Find("aside,div.social-share,div.injected-related-story,div.a-article-after-content").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})

	document.Find("div.a-article-grid__main").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) HollywoodreporterExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := hollywoodreporterScrapContent(document)
	return content, "", 0, "", "", ""
}
