package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func ignScrapContent(document *goquery.Document) string {
	contents := ""
	document.Find("section.user-list-embed").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})
	document.Find("img").Each(func(i int, img *goquery.Selection) {
		parent := img.Parent()
		parentNode := parent.Get(0)
		if parentNode.Data == "a" {
			href, exists := parent.Attr("href")
			if exists {
				img.SetAttr("src", href)
			}
		}
	})
	document.Find("div.article-header-image,section.article-page").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) IGNExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := ignScrapContent(document)
	return content, "", 0, "", "", ""
}
