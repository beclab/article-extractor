package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func scmpScrapContent(document *goquery.Document) string {
	contents := ""
	document.Find("div[data-qa=GenericArticle-Leading],div[data-qa=GenericArticle-Content]").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents

}

func (t *Template) SCMPExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := scmpScrapContent(document)
	return content, "", 0, "", "", ""
}
