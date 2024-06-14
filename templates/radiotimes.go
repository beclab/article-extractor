package templates

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func (t *Template) RadiotimesScrapContent(document *goquery.Document) string {
	contents := ""

	document.Find("div.ad-placement,div[data-feature=NextRead],div.ad-slot,div.newsletter-sign-up,div.rating").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})
	document.Find("p,strong").Each(func(i int, s *goquery.Selection) {
		text := s.Text()

		if strings.HasPrefix(text, "Read more:") {
			n := s.Next()
			if n != nil && len(n.Children().Nodes) > 1 {
				htmlNode := n.Get(0)
				if htmlNode.Data == "ul" {
					RemoveNodes(n)
				}
			}

			RemoveNodes(s)

		}

	})
	document.Find("div.post-header__image-container,div.post__content").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}
