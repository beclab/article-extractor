package templates

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func (t *Template) RumbleExtractorMetaInfo(entryUrl string, document *goquery.Document) (string, string, int64, string, string, string) {

	embeddingUrl := ""
	document.Find("link[type='application/json+oembed']").Each(func(i int, s *goquery.Selection) {
		if href, exists := s.Attr("href"); exists {
			index := strings.Index(href, "=")
			if index != -1 {
				embeddingUrl = href[index+1:]
				decodedString, err := url.QueryUnescape(embeddingUrl)
				if err != nil {
					fmt.Println("runble url decode error :", err)
					return
				} else {
					embeddingUrl = decodedString
				}
			}
		}

	})

	if embeddingUrl != "" {
		contents := "<iframe width='960' height='540' src='" + embeddingUrl + "'  frameborder='0'  referrerpolicy='no-referrer'></iframe>"
		return "", "", 0, contents, entryUrl, "video"
	}
	return "", "", 0, "", "", ""
}
