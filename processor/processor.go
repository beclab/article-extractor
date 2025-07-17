package processor

import (
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/beclab/article-extractor/readability"
	"github.com/beclab/article-extractor/rewrite"
	"github.com/beclab/article-extractor/sanitizer"
	"github.com/beclab/article-extractor/templates"
	"github.com/beclab/article-extractor/templates/postExtractor"
)

func ArticleContentExtractor(rawContent, entryUrl, feedUrl, rules string) (string, string) {
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
		postContent := applyPostExtraction(contentPreRule, content, feedUrl)
		if postContent != "" {
			content = postContent
		}
	}

	pureContent := getPureContent(content)
	return content, pureContent
}

func NonRawConntentDownloadQueryInArticle(url string) (string, string, string) {
	funcs := reflect.ValueOf(&templates.Template{})
	_, mediaRule := getNonRawContentDownloadScraperRules(url)
	if mediaRule != "" {
		f := funcs.MethodByName(mediaRule)
		res := f.Call([]reflect.Value{reflect.ValueOf(url)})
		downloadUrl := res[0].String()
		downloadFile := res[1].String()
		downloadType := res[2].String()
		if downloadType == "audio" || downloadType == "ebook" || downloadType == "pdf" {
			return downloadUrl, downloadFile, downloadType
		}
	}
	return "", "", ""
}

func ArticleReadabilityExtractor(rawContent, entryUrl, feedUrl, rules string, isrecommend bool) (string, string, *time.Time, string, string, string, int64, string, string, string) {
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

	var content string
	var author string
	var mediaContent string
	var mediaUrl string
	var mediaType string
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
		mediaUrl = res[4].String()
		mediaType = res[5].String()
	}
	if author == "" {
		if strings.Contains(entryUrl, "weixin.qq.com") {
			author = GetAuthorForWechat(doc)
		} else {
			author = templates.ScrapAuthorMetaData(doc)
		}
	}
	if publishedAt == 0 {
		if strings.Contains(entryUrl, "weixin.qq.com") {
			publishedAt = GetPublishedAtTimestampForWechat(rawContent, entryUrl)
		} else {
			publishedAt = templates.ScrapPublishedAtTimeMetaData(doc)
		}
	}

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
		postContent := applyPostExtraction(contentPreRule, article.Content, feedUrl)
		if postContent != "" {
			article.Content = postContent
		}
	}

	pureContent := getPureContent(article.Content)
	updateTitle := updateArticleTitle(entryDomain, doc)
	if updateTitle != "" {
		article.Title = updateTitle
	}
	return article.Content, pureContent, article.PublishedDate, article.Image, article.Title, author, publishedAt, mediaContent, mediaUrl, mediaType
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
func applyPostExtraction(contentPreRule, articleContent, feedUrl string) string {
	postFuncs := reflect.ValueOf(&postExtractor.PostExtractorTemplate{})
	f := postFuncs.MethodByName(contentPreRule)
	res := f.Call([]reflect.Value{reflect.ValueOf(articleContent), reflect.ValueOf(feedUrl)})
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

func GetPublishedAtTimestampForWechat(rawContent string, url string) int64 {
	var publishedAtTimestamp int64 = 0
	re := regexp.MustCompile(`var oriCreateTime = '(\d+)';`)
	match := re.FindStringSubmatch(rawContent)
	if len(match) > 1 {
		timestamp, err := strconv.ParseInt(match[1], 10, 64)
		if err != nil {
			log.Printf("can not parse timestamp [%s] for entry [%s]", match[1], url)
			return publishedAtTimestamp
		}
		publishedAtTimestamp = timestamp
	} else {
		log.Printf("can not find timestamp for entry [%s]", url)
		return publishedAtTimestamp
	}
	return publishedAtTimestamp
}

func GetAuthorForWechat(document *goquery.Document) string {
	var author string
	// Function to extract author from a given selector
	extractAuthor := func(selector string) {
		document.Find(selector).Each(func(i int, s *goquery.Selection) {
			content := strings.TrimSpace(s.Text())
			content = strings.ReplaceAll(content, "\n", "")
			author = content
		})
	}
	// Try to extract author from both selectors
	extractAuthor("div#meta_content>span.rich_media_meta_text")
	if author == "" {
		extractAuthor("div#meta_content>span.rich_media_meta_nickname")
	}
	return author
}
