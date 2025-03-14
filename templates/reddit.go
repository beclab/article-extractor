package templates

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/beclab/article-extractor/readability"
)

func (t *Template) RedditScrapContent(document *goquery.Document) string {
	contents := ""

	document.Find("div[slot='text-body'],img#post-image").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})

	/*document.Find("shreddit-player-2").Each(func(i int, s *goquery.Selection) {
		contents = contents + "<video>"
		s.Find("source").Each(func(i int, sourceSel *goquery.Selection) {
			var content string
			content, _ = goquery.OuterHtml(sourceSel)
			contents += content
		})
		contents = contents + "</video>"
	})*/

	return contents
}

func (t *Template) RedditMediaContent(url string, document *goquery.Document) (string, string, string) {
	videoUrl := ""
	videoType := ""
	document.Find("shreddit-player-2>source").Each(func(i int, s *goquery.Selection) {
		videoUrl = url
		videoType = "video"
	})

	return "", videoUrl, videoType

}

func (t *Template) RedditScrapMetaData(document *goquery.Document) (string, string) {

	author := ""
	published_at := ""

	document.Find("shreddit-post").Each(func(i int, s *goquery.Selection) {
		if content, exists := s.Attr("author"); exists {
			author = content
		}
	})
	return author, published_at
}

func (t *Template) RedditPublishedAtTimeFromScriptMetadata(document *goquery.Document) int64 {

	var publishedAt int64 = 0

	document.Find("shreddit-post").Each(func(i int, s *goquery.Selection) {
		if publishTimes, exists := s.Attr("created-timestamp"); exists {
			dateObj, err := readability.ParseTime(publishTimes)
			if err == nil {
				publishedAt = dateObj.Unix()
			}
		}
	})

	return publishedAt
}
