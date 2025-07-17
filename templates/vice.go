package templates

import (
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

func viceContentExtractor(document *goquery.Document) string {
	contents := ""
	document.Find("div.adph,div.abc__article_embed").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})
	//https://www.vice.com/en/article/wxjymy/i-rode-melbournes-free-taylor-swift-trams-for-12-hours-because-i-love-free-things-and-hate-myself
	/*
			 <picture class="responsive-image lazyloader--lazy lazyloader--lowres"><source media="(min-width: 1000px)" srcSet="https://video-images.vice.com/_uncategorized/1708388089098-img0972.jpeg?resize=20:*"/>
		          <source media="(min-width: 700px)" srcSet="https://video-images.vice.com/_uncategorized/1708388089098-img0972.jpeg?resize=20:*"/>
		          <source media="(min-width: 0px)" srcSet="https://video-images.vice.com/_uncategorized/1708388089098-img0972.jpeg?resize=20:*"/>
		          <img class="responsive-image__img" alt="man on tram" decoding="async" loading="eager" width="2316" height="1745"/></picture>
	*/
	var rxPicture = regexp.MustCompile(`resize=(\d+):`)
	document.Find("picture").Each(func(i int, pic *goquery.Selection) {
		firstChild := pic.Children().First()
		childNode := firstChild.Get(0)
		if childNode.Data == "source" {
			src, exists := firstChild.Attr("srcset")
			if exists {
				replaceVal := rxPicture.ReplaceAllString(src, "resize=650:")
				firstChild.SetAttr("srcset", replaceVal)
			}
		}
	})

	document.Find("div.short-form__body__article-lede-image,div.article__body-components").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})
	return contents
}

func (t *Template) ViceExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := viceContentExtractor(document)
	return content, "", 0, "", "", ""
}
