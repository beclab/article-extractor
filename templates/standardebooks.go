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

func extractStandardebooksName(urlStr string) string {
	lastSlashIndex := strings.LastIndex(urlStr, "/")
	fileName := urlStr[lastSlashIndex+1:]
	return fileName
}

func (t *Template) StandardebooksNonRawContent(url string, document *goquery.Document) (string, string, string) {
	if isEpubURL(url) {
		fileName := extractStandardebooksName(url)
		url = url + "?source=download"
		return url, fileName, "ebook"
	}
	return "", "", ""

}
