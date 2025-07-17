package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func deadlineScrapContent(document *goquery.Document) string {
	contents := ""
	document.Find("div.injected-related-story,section.toaster,div.article-tags,div#comments-loading,h2#comments-title,p.subscribe-to").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})
	document.Find("article").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) DeadlineExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := deadlineScrapContent(document)

	return content, "", 0, "", "", ""
}
