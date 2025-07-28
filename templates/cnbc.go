package templates

import (
	"github.com/PuerkitoBio/goquery"
)

func cnbcScrapContent(document *goquery.Document) string {
	contents := ""
	document.Find("div#RegularArticle-RelatedQuotes,div[data-test=PlayButton],div.InlineVideo-videoFooter,div.InlineImage-imageEmbedCaption,div.InlineImage-imageEmbedCredit,div.RelatedContent-relatedContent").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})

	document.Find("div.InlineImage-imageContainer,div.RenderKeyPoints-list,div.ArticleBody-articleBody").Each(func(i int, s *goquery.Selection) {
		var content string
		content, _ = goquery.OuterHtml(s)
		contents += content
	})

	return contents
}

func (t *Template) CNBCExtractorMetaInfo(url string, document *goquery.Document) (string, string, int64, string, string, string) {
	content := cnbcScrapContent(document)
	return content, "", 0, "", "", ""
}
