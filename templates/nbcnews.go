package templates

import (
	"github.com/PuerkitoBio/goquery"
)

// site url https://www.nbcnews.com/world
// rss  http://feeds.nbcnews.com/feeds/worldnews

func nbcNewsScrapContent(document *goquery.Document) string {
	contents := ""
	document.Find("h1.article-hero-headline__htag,aside").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})
	document.Find("section.article-hero__container").Each(func(i int, s *goquery.Selection) {
		content, _ := goquery.OuterHtml(s)
		contents += content
	})
	articleSectionBody := document.Find("div.article-body__first-section").First()
	articleBody := articleSectionBody.Find("div").First()
	firstArticle := articleBody.Find("div.article-body__content").First()

	articleContent, _ := goquery.OuterHtml(firstArticle)
	contents += articleContent

	return contents
}
func (t *Template) NbcNewsExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := nbcNewsScrapContent(document)
	return content, "", 0, "", "", ""
}
