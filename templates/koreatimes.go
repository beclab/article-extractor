package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func (t *Template) KoreatimesScrapMetaData(document *goquery.Document) (string, string) {

	author := "KoreaTimes"
	published_at := ""

	return author, published_at
}
