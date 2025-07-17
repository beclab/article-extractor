package templates

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/beclab/article-extractor/readability"
)

func redditScrapContent(document *goquery.Document) string {
	contents := ""
	document.Find("div[slot='text-body'],img#post-image").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})

	return contents
}

func (t *Template) RedditExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := redditScrapContent(document)
	author := ""
	videoUrl := ""
	var publishedAt int64 = 0
	document.Find("shreddit-post").Each(func(i int, s *goquery.Selection) {
		if content, exists := s.Attr("author"); exists {
			author = content
		}
	})
	document.Find("shreddit-post").Each(func(i int, s *goquery.Selection) {
		if publishTimes, exists := s.Attr("created-timestamp"); exists {
			dateObj, err := readability.ParseTime(publishTimes)
			if err == nil {
				publishedAt = dateObj.Unix()
			}
		}
	})
	document.Find("shreddit-player-2>source").Each(func(i int, s *goquery.Selection) {
		videoUrl = url
	})
	return content, author, publishedAt, "", videoUrl, ""
}
