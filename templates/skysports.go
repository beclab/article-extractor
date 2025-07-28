package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func skySportsScrapContent(document *goquery.Document) string {

	contents := ""
	document.Find("div.sdc-article-related-stories,div.sdc-article-factbox,div.sdc-article-strapline,div.site-share-wrapper,div[data-format=floated-mpu]").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})

	document.Find("div.sdc-article-image,div[data-type=article]").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents

}

func (t *Template) SkySportsExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := skySportsScrapContent(document)
	return content, "", 0, "", "", ""
}
