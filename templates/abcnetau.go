package templates

import "github.com/PuerkitoBio/goquery"

func abcNetAUScrapContent(document *goquery.Document) string {
	contents := ""
	document.Find("aside,div.Headline_meta__ZgyGe,div[data-component=RelatedTopics],div[data-component=ShareUtility],div[data-component=Dateline]").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})

	document.Find("div.Article_layoutMain__eBEMA").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) AbcNetExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := abcNetAUScrapContent(document)
	author := "abc.net.au"
	return content, author, 0, "", "", ""
}
