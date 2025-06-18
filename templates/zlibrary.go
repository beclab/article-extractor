package templates

import (
	"fmt"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

func (t *Template) ZLibraryMediaContent(url string, document *goquery.Document) (string, string, string) {
	downloadUrl := ""
	pattern := `^https:\/\/z-library\.gs\/book\/.*`
	matched, err := regexp.MatchString(pattern, url)
	if err != nil {
		fmt.Println("zlib match err:", err)
		return "", "", ""
	}
	if matched {
		document.Find("a.addDownloadedBook").Each(func(i int, s *goquery.Selection) {
			href, exists := s.Attr("href")
			if exists {
				downloadUrl = "https://z-library.gs" + href
				return
			}
		})
	}
	if downloadUrl != "" {
		return downloadUrl, downloadUrl, "ebook"
	}
	pattern = `^https:\/\/z-library\.gs\/dl\/.*`
	matched, err = regexp.MatchString(pattern, url)
	if err != nil {
		fmt.Println("zlib match err2:", err)
		return "", "", ""
	}
	if matched {
		return url, url, "ebook"
	}

	return "", "", ""

}
