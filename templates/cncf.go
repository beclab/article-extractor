package templates

import (
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func cncfScrapContent(document *goquery.Document) string {
	contents := ""
	document.Find("div.social-share,figure.wp-block-embed-twitter,div.post-author").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})

	document.Find("article.container").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func cncfScrapAuthor(document *goquery.Document) string {
	author := ""
	document.Find("span.post-author__author").Each(func(i int, s *goquery.Selection) {
		spanContent := s.Text()
		author = strings.TrimPrefix(spanContent, "By ")
	})
	if len(author) == 0 {
		document.Find("div.post-author").Next().Each(func(i int, s *goquery.Selection) {
			emContent := s.Find("em").Text()
			cleanContent := strings.Replace(emContent, "\u00a0", "", -1)
			cleanContent = strings.TrimPrefix(cleanContent, "Community post by ")
			cleanContent = strings.TrimPrefix(cleanContent, "Member post by ")
			cleanContent = strings.TrimPrefix(cleanContent, "Community post originally published on Medium by ")
			cleanContent = strings.TrimPrefix(cleanContent, "Project post by ")
			cleanContent = strings.TrimPrefix(cleanContent, "Member post originally published on Greptime’s blog by ")
			extractAuthor, extractSuccess := extractAuthorFromCNCFName(cleanContent)
			if extractSuccess {
				author = extractAuthor
			} else {
				author = cleanContent
			}
		})

	}

	return author
}

func extractAuthorFromCNCFName(input string) (string, bool) {
	re := regexp.MustCompile(`Member post originally published on ([^’]+)’s blog`)
	matches := re.FindStringSubmatch(input)
	if len(matches) >= 2 {
		return matches[1], true
	}
	return "", false
}

func (t *Template) CNCFExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := cncfScrapContent(document)
	author := cncfScrapAuthor(document)
	return content, author, 0, "", "", ""
}
