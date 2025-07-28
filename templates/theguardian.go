package templates

import (
	"github.com/PuerkitoBio/goquery"
)

type TheGuardianCoverImage struct {
	TheImageUrlList []string `json:"image"`
}

func theguardianScrapContent(document *goquery.Document) string {
	contents := ""
	document.Find("gu-island,aside,a.dcr-porppu,p.dcr-porppu,p#EmailSignup-skip-link-8").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})
	document.Find("a").Each(func(i int, s *goquery.Selection) {
		if s.Text() == "skip past newsletter promotion" {
			RemoveNodes(s)
		}

	})
	document.Find("div#img-1,div#maincontent").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) TheguardianExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := theguardianScrapContent(document)
	return content, "", 0, "", "", ""
}
