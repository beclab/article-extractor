package templates

import (
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type InterestingengineeringMetaData struct {
	Context string `json:"@context"`
	Graph   []struct {
		Type   string `json:"@type"`
		Author struct {
			Name string `json:"name"`
			ID   string `json:"@id"`
		} `json:"author"`
		DatePublished time.Time `json:"datePublished"`
		DateModified  time.Time `json:"dateModified"`
	} `json:"@graph"`
}

func (t *Template) InterestingengineeringPublishedAtTimeFromScriptMetadata(doc *goquery.Document) int64 {
	var publishedAt int64 = 0

	scriptSelector := "script[type=\"application/ld+json\"]"
	doc.Find(scriptSelector).Each(func(i int, s *goquery.Selection) {
		scriptContent := strings.TrimSpace(s.Text())
		var metaData InterestingengineeringMetaData
		unmarshalErr := json.Unmarshal([]byte(scriptContent), &metaData)
		if unmarshalErr != nil {
			log.Printf("convert SkyNewsScrap unmarshalError %v", unmarshalErr)
		} else {
			for _, graphData := range metaData.Graph {
				publishedAt = graphData.DatePublished.Unix()
				return
			}
		}

	})

	return publishedAt
}

func (t *Template) InterestingengineeringScrapContent(document *goquery.Document) string {
	contents := ""

	document.Find("div.social-icons,div.article-heading,div#articleRight,div.Ad_adContainer__XNCwI,div[data-orientation=vertical],div.article-thumbnail--info,div.SubscriptionInlineForm_newsletterContainer__HotUe,div.recommendedArticle_recommended_article__ENN1_,div.CommentSection_commentsblock__cerVm,nav").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)

	})
	document.Find("div.article-related--items").Each(func(i int, s *goquery.Selection) {
		p := s.Parent()
		if p.Parent() != nil {
			RemoveNodes(p.Parent())
		} else {
			RemoveNodes(p)
		}
	})

	document.Find("article").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}
