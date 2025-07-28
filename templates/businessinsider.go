package templates

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func businessInsiderScrapContent(document *goquery.Document) string {
	contents := ""
	document.Find("label.caption-drawer-label,div.piano-inline-content-wrapper,div.in-post-sticky,div.inline-newsletter-signup,div.ad-callout-wrapper,article.d-none,section.content-recommendations-component,section.related-posts").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})
	document.Find("img").Each(func(i int, img *goquery.Selection) {
		srcValid := false
		if srcAttr, found := img.Attr("src"); found {
			if strings.HasPrefix(srcAttr, "http") {
				srcValid = true
			}
		}
		if !srcValid {
			if datasrcsAttr, found := img.Attr("data-srcs"); found {
				var dataResult map[string]map[string]interface{}

				err := json.Unmarshal([]byte(datasrcsAttr), &dataResult)
				if err != nil {
					fmt.Println("data-srcs to json error:", err)
				}
				for url := range dataResult {
					img.SetAttr("src", url)
					break
				}
			}
		}
	})
	document.Find("div.post-hero,div.post-summary-bullets,section.post-body-content").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) BusinessInsiderExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := businessInsiderScrapContent(document)
	return content, "", 0, "", "", ""
}
