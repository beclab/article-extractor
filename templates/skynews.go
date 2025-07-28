package templates

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func skyNewsScrapContent(document *goquery.Document) string {
	contents := ""
	document.Find("div.sdc-article-related-stories,div.sdc-site-video,a,span[data-label-text=Advertisement],div.sdc-article-related-stories,div.sdc-article-strapline").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})
	document.Find("p").Each(func(i int, s *goquery.Selection) {
		if strings.Contains(s.Text(), "Read more:") {
			RemoveNodes(s)
		}
	})
	document.Find("div.sdc-site-component-top__media,div.sdc-article-body").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) SkyNewsExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := skyNewsScrapContent(document)
	return content, "", 0, "", "", ""
}
