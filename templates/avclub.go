package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func avClubScrapContent(document *goquery.Document) string {
	contents := ""
	document.Find("div.instream-native-video,div.advertisement,div.js_ad-mobile-dynamic,div.js_related-stories-inset,div.js_related-stories-inset-mobile,div.js_postbottom-waypoint-hook,div.js_comments-iframe").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})

	document.Find("div.article-header__feature,div.js_post-content").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) AVClubExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := avClubScrapContent(document)
	return content, "", 0, "", "", ""
}
