package templates

import (
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

func (t *Template) YoutubeScrapContent(document *goquery.Document) string {
	contents := ""
	/*document.Find("#description-inline-expander").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})*/
	return contents
}

func (t *Template) YoutubeMediaContent(url string, document *goquery.Document) (string, string, string) {
	//pattern := `^https?://(?:www\.)?youtube\.com/watch\?v=([a-zA-Z0-9_-]+).*$`
	pattern := `youtube\.com/watch\?v=([^&]+)`
	regex := regexp.MustCompile(pattern)
	match := regex.FindStringSubmatch(url)
	if match != nil {
		if len(match) > 1 {
			videoID := match[1]
			embedUrl := "https://www.youtube.com/embed/gfx7mTmWdYU?si=" + videoID
			contents := "<iframe width='840' height='472' src='" + embedUrl + "'  frameborder='0' allow='accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share' referrerpolicy='strict-origin-when-cross-origin' allowfullscreen></iframe>"
			return contents, url, "video"
		}

	}
	return "", "", ""

}

func (t *Template) YoutubeScrapMetaData(document *goquery.Document) (string, string) {

	author := ""
	published_at := ""
	/*scriptSelector := "script[type=\"application/ld+json\"]"
	document.Find(scriptSelector).Each(func(i int, s *goquery.Selection) {
		scriptContent := strings.TrimSpace(s.Text())
		var metaData map[string]interface{}
		unmarshalErr := json.Unmarshal([]byte(scriptContent), &metaData)
		if unmarshalErr != nil {
			log.Printf("convert  unmarshalError %v", unmarshalErr)
		}
		if _, ok := metaData["author"]; ok {
			author = metaData["author"].(string)
			return
		}

	})*/

	return author, published_at
}

func (t *Template) YoutubePublishedAtTimeFromScriptMetadata(doc *goquery.Document) int64 {
	var publishedAt int64 = 0

	/*s := doc.Find("meta[itemprop='datePublished']").First()
	timeStr, exists := s.Attr("content")
	if exists {
		ptime, parseErr := readability.ParseTime(timeStr)
		if parseErr == nil {
			return ptime.Unix()
		}
	}*/

	return publishedAt
}
