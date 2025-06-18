package templates

import (
	"fmt"
	"net/url"
	"path"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func extractIDWithRegex(urlStr string) string {
	re := regexp.MustCompile(`/get/([^/]+)/[26]$`)
	matches := re.FindStringSubmatch(urlStr)

	if len(matches) < 2 {
		return ""
	}

	return matches[1]
}

func (t *Template) ManyBooksMediaContent(urlStr string, document *goquery.Document) (string, string, string) {
	u, err := url.Parse(urlStr)
	if err != nil {
		fmt.Println("many books parse err:", err)
		return "", "", ""
	}

	cleanPath := strings.Trim(u.Path, "/")
	lastPart := path.Base(cleanPath)
	if lastPart == "2" {
		//ebook
		id := extractIDWithRegex(urlStr)
		downurl := "https://library.manybooks.net/live/get-book/" + id + "/epub"
		return downurl, downurl, "ebook"
	}
	if lastPart == "6" {
		//pdf
		id := extractIDWithRegex(urlStr)
		downurl := "https://library.manybooks.net/live/get-book/" + id + "/pdf"
		return downurl, downurl, "pdf"

	}

	return "", "", ""
}
