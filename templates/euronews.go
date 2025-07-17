package templates

import (
	"github.com/PuerkitoBio/goquery"
)

// site url  https://www.euronews.com/news/international
// rss  https://www.euronews.com/rss?format=mrss&level=theme&name=news

func euroNewsScrapContent(document *goquery.Document) string {
	contents := ""
	document.Find("nav,h1.c-article-redesign-title,div.c-article-you-might-also-like,div.c-article-contributors,time.c-article-publication-date,div.c-ad__placeholder,a.c-article-partage-commentaire__links,div.c-article-caption,div.c-article-partage-commentaire-popup-overlay").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})
	document.Find("div.o-article-newsy__main").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) EuroNewsExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := euroNewsScrapContent(document)

	return content, "", 0, "", "", ""
}
