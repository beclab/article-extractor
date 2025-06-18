package templates

import (
	"net/url"
	"path"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func isEpubURL(rawURL string) bool {
	u, err := url.Parse(rawURL)
	if err != nil {
		return false
	}

	filename := path.Base(u.Path)
	if filename == "." || filename == "/" {
		return false
	}

	ext := strings.ToLower(path.Ext(filename))
	return ext == ".epub"
}

func (t *Template) StandardebooksMediaContent(url string, document *goquery.Document) (string, string, string) {
	if isEpubURL(url) {
		return url, url, "ebook"
	}
	downloadUrl := ""
	document.Find("a.pub").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			downloadUrl = "https://standardebooks.org/" + href
			return
		}
	})
	return downloadUrl, downloadUrl, "ebook"
}
