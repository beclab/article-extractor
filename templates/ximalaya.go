package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func ximalayaScrapContent(document *goquery.Document) string {
	contents := ""

	document.Find("article.intro").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) XimalayaExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := ximalayaScrapContent(document)
	author := ""
	fileType := ""
	document.Find("a.albumTitle").Each(func(i int, s *goquery.Selection) {
		author = s.Text()
		fileType = AudioFileType
	})
	return content, author, 0, url, url, fileType
}
