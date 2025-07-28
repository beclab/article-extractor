package processor

import (
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/beclab/article-extractor/readability"
	"github.com/beclab/article-extractor/rewrite"
	"github.com/beclab/article-extractor/sanitizer"
	"github.com/beclab/article-extractor/templates"
	"github.com/beclab/article-extractor/templates/postExtractor"
)

// 得到content内容，主要在推荐算法爬取页面后解析正文内容
func ArticleContentExtractor(rawContent, entryUrl string) (string, string) {
	entryDomain := domain(entryUrl)
	templateRawData := strings.NewReader(rawContent)
	doc, _ := goquery.NewDocumentFromReader(templateRawData)

	var content string
	funcs := reflect.ValueOf(&templates.Template{})
	contentRule := getPredefinedRules(entryUrl, doc)
	if contentRule != "" {
		f := funcs.MethodByName(contentRule)
		res := f.Call([]reflect.Value{reflect.ValueOf(entryUrl), reflect.ValueOf(doc)})
		content = res[0].String()
	}

	if content != "" {
		content = processContent(content, entryDomain, entryUrl)
	} else {
		content = templates.GetArticleByDivClass(doc)
	}

	if content == "" {
		rawData := strings.NewReader(rawContent)
		article, err := readability.FromReader(rawData, entryUrl)
		if err != nil {
			log.Printf(`article extractor error %q`, err)
			return "", ""
		}
		content = article.Content
	}
	contentPreRule := getContentPostExtractorTemplateRules(entryUrl)
	if content != "" && contentPreRule != "" {
		postContent := applyPostExtraction(contentPreRule, content)
		if postContent != "" {
			content = postContent
		}
	}

	pureContent := getPureContent(content)
	return content, pureContent
}

// 根据url，不用正文内容获得下载信息
// 对于ebook和pdf 通过url来解析，不需要爬取页面
func DownloadTypeQueryByUrl(url string) (string, string, string) {
	funcs := reflect.ValueOf(&templates.Template{})
	_, mediaRule := getDownloadTypeByUrlRules(url)
	if mediaRule != "" {
		f := funcs.MethodByName(mediaRule)
		res := f.Call([]reflect.Value{reflect.ValueOf(url)})
		downloadUrl := res[0].String()
		downloadFile := res[1].String()
		downloadType := res[2].String()
		return downloadUrl, downloadFile, downloadType
	}
	return "", "", ""
}

// 根据模版表获得正文,作者，发布时间，以及下载信息
func MetaDataQueryByTemplate(entryUrl, rawContent string, doc *goquery.Document) (string, string, int64, string, string, string) {
	var content string
	var author string
	var mediaContent string
	var downloadUrl string
	var downloadType string
	var publishedAt int64 = 0
	funcs := reflect.ValueOf(&templates.Template{})
	contentRule := getPredefinedRules(entryUrl, doc)
	if contentRule != "" {
		f := funcs.MethodByName(contentRule)
		res := f.Call([]reflect.Value{reflect.ValueOf(entryUrl), reflect.ValueOf(doc)})
		content = res[0].String()
		author = res[1].String()
		publishedAt = res[2].Int()
		mediaContent = res[3].String()
		downloadUrl = res[4].String()
		downloadType = res[5].String()
	}
	if author == "" {
		if strings.Contains(entryUrl, "weixin.qq.com") {
			author = templates.GetAuthorForWechat(doc)
		} else {
			author = templates.ScrapAuthorMetaData(doc)
		}
	}
	if publishedAt == 0 {
		if strings.Contains(entryUrl, "weixin.qq.com") {
			publishedAt = templates.GetPublishedAtTimestampForWechat(rawContent, entryUrl)
		} else {
			publishedAt = templates.ScrapPublishedAtTimeMetaData(doc)
		}
	}
	return content, author, publishedAt, mediaContent, downloadUrl, downloadType
}

// 输入url，rawcontent
// 输出entry的metadata
func ArticleExtractor(rawContent, entryUrl string) (string, string, *time.Time, string, string, string, int64, string, string, string) {
	templateRawData := strings.NewReader(rawContent)
	doc, _ := goquery.NewDocumentFromReader(templateRawData)

	entryDomain := domain(entryUrl)
	rawData := strings.NewReader(rawContent)
	article, err := readability.FromReader(rawData, entryUrl)
	log.Printf("get readability article %s", entryUrl)
	if err != nil {
		log.Printf(`article extractor error %q`, err)
		return "", "", nil, "", "", "", 0, "", "", ""
	}

	content, author, publishedAt, mediaContent, downloadUrl, downloadType := MetaDataQueryByTemplate(entryUrl, rawContent, doc)
	if content != "" {
		content = processContent(content, entryDomain, entryUrl)
	} else {
		content = templates.GetArticleByDivClass(doc)
	}
	if content != "" {
		article.Content = content
	}

	contentPreRule := getContentPostExtractorTemplateRules(entryUrl)
	if article.Content != "" && contentPreRule != "" {
		postContent := applyPostExtraction(contentPreRule, article.Content)
		if postContent != "" {
			article.Content = postContent
		}
	}

	pureContent := getPureContent(article.Content)
	updateTitle := updateArticleTitle(entryDomain, doc)
	if updateTitle != "" {
		article.Title = updateTitle
	}
	return article.Content, pureContent, article.PublishedDate, article.Image, article.Title, author, publishedAt, mediaContent, downloadUrl, downloadType
}

func updateArticleTitle(entryDomain string, doc *goquery.Document) string {
	updateTitle := ""
	if strings.Contains(entryDomain, "reddit.com") {
		doc.Find("shreddit-post").Each(func(i int, s *goquery.Selection) {
			if title, exists := s.Attr("post-title"); exists {
				updateTitle = strings.TrimSpace(title)
			}
		})
	}
	return updateTitle
}
func applyPostExtraction(contentPreRule, articleContent string) string {
	postFuncs := reflect.ValueOf(&postExtractor.PostExtractorTemplate{})
	f := postFuncs.MethodByName(contentPreRule)
	res := f.Call([]reflect.Value{reflect.ValueOf(articleContent)})
	return res[0].String()
}

func isSanitizeRequired(entryDomain string) bool {
	nonSanitizeDomains := []string{"okjike.com", "vimeo.com", "fandom.com", "notion.site", "quora.com"}
	for _, domain := range nonSanitizeDomains {
		if strings.Contains(entryDomain, domain) {
			return false
		}
	}
	return true
}

func processContent(content, entryDomain, entryUrl string) string {
	if !strings.Contains(entryDomain, "douban.com") {
		content = rewrite.Rewriter(entryUrl, content, "add_dynamic_image")
	}
	if isSanitizeRequired(entryDomain) {
		return sanitizer.Sanitize(entryUrl, content)
	}
	return content
}

func getPureContent(content string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err == nil {
		pureText := doc.Text()
		pureText = strings.Replace(pureText, "\n", "", -1)
		pureText = strings.Replace(pureText, "\t", "", -1)
		return pureText
	}
	return ""
}
