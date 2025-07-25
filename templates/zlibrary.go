package templates

import (
	"fmt"
	"regexp"
)

func extractZLibraryIDWithRegex(urlStr string) string {
	re := regexp.MustCompile(`/dl/(\d+)/`)
	matches := re.FindStringSubmatch(urlStr)

	if len(matches) < 2 {
		return ""
	}

	return matches[1]
}

/*func (t *Template) ZLibraryDownloadFromWeb(url string, document *goquery.Document) []model.ExtractorFileInfo {
	downloadUrl := ""
	var fileList []model.ExtractorFileInfo
	pattern := `^https:\/\/z-library\.gs\/book\/.*`
	matched, err := regexp.MatchString(pattern, url)
	if err != nil {
		fmt.Println("zlib match err:", err)
		return fileList
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
		fileList = append(fileList, model.ExtractorFileInfo{DownloadUrl: downloadUrl, FileName: extractZLibraryIDWithRegex(downloadUrl) + ".epub", FileType: "ebook"})
		return fileList
	}
	return fileList
}*/

func (t *Template) ZLibraryDownloadType(url string) (string, string, string) {
	pattern := `^https:\/\/z-library\.gs\/dl\/.*`
	matched, err := regexp.MatchString(pattern, url)
	if err != nil {
		fmt.Println("zlib match err2:", err)
		return "", "", ""
	}
	if matched {
		return url, extractZLibraryIDWithRegex(url) + ".epub", "ebook"
	}
	return "", "", ""

}
