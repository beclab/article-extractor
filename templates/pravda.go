package templates

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// site url  https://www.pravda.com.ua/eng/news/
// rss  https://www.pravda.com.ua/eng/rss/view_news/
func pravdaScrapAuthor(document *goquery.Document) string {
	author := ""
	cssSelectorFirst := "body > div.main_content > div > div.container_sub_post_news > article > header > div > span > a"

	cssSelectorList := make([]string, 100)
	cssSelectorList = append(cssSelectorList, cssSelectorFirst)

	for _, cssSelector := range cssSelectorList {
		document.Find(cssSelector).Each(func(i int, s *goquery.Selection) {
			author = strings.TrimSpace(s.Text())
		})
		if author != "" {
			break
		}
	}
	return author
}

func pravdaScrapContent(document *goquery.Document) string {
	contents := ""
	document.Find("div.advtext_mob").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})

	document.Find("div.post_photo_news,div.post_text").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) PravdaExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := pravdaScrapContent(document)
	author := pravdaScrapAuthor(document)
	return content, author, 0, "", "", ""
}
