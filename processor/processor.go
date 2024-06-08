package processor

import (
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/beclab/article-extractor/readability"
	"github.com/beclab/article-extractor/rewrite"
	"github.com/beclab/article-extractor/sanitizer"
)

func ArticleContentExtractor(pluginsPath, rawContent, entryUrl, feedUrl, rules string) (string, string) {
	templateRawData := strings.NewReader(rawContent)
	doc, _ := goquery.NewDocumentFromReader(templateRawData)

	content := getPluginsContent(pluginsPath, entryUrl, doc)
	if content == "" {
		var ruleErr error
		var rulesDomain string
		contentRule := rules
		if contentRule == "" {
			rulesDomain, contentRule = getPredefinedScraperRules(entryUrl)
		}
		if contentRule != "" {
			content, ruleErr = ScrapContentUseRules(doc, rules)
			if ruleErr != nil {
				log.Printf(`get document by rule error rules:%s,domain:%s,%q`, rules, rulesDomain, ruleErr)
				return "", ""
			}
		}
	}
	if content != "" {
		content = rewrite.Rewriter(entryUrl, content, "add_dynamic_image")
		content = sanitizer.Sanitize(entryUrl, content)
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

	postContent := getPluginsPostContent(pluginsPath, entryUrl, content)
	if postContent != "" {
		content = postContent
	}
	pureContent := getPureContent(content)
	return content, pureContent
}

func ArticleReadabilityExtractor(pluginsPath, rawContent, entryUrl, feedUrl, rules string, isrecommend bool) (string, string, *time.Time, string, string, string, string, int64) {
	templateRawData := strings.NewReader(rawContent)
	doc, _ := goquery.NewDocumentFromReader(templateRawData)

	rawData := strings.NewReader(rawContent)
	article, err := readability.FromReader(rawData, entryUrl)
	log.Printf("get readability article %s", entryUrl)
	if err != nil {
		log.Printf(`article extractor error %q`, err)
		return "", "", nil, "", "", "", "", 0
	}

	author := getPluginsAuthor(pluginsPath, entryUrl, doc)
	if author == "" && !strings.HasPrefix(feedUrl, "wechat") {
		author = ScrapAuthorMetaData(doc)
	}

	publishedAtTimeStamp := getPluginsPublishedAtTemplate(pluginsPath, entryUrl, doc)
	if publishedAtTimeStamp == 0 && !strings.HasPrefix(feedUrl, "wechat") {
		publishedAtTimeStamp = ScrapAutoPublishedAtTimeMetaData(doc)
	}
	if strings.HasPrefix(feedUrl, "wechat") {
		publishedAtTimeStamp = GetPublishedAtTimestampForWechat(rawContent, entryUrl)
	}

	content := getPluginsContent(pluginsPath, entryUrl, doc)
	if content == "" {
		var ruleErr error
		var rulesDomain string
		contentRule := rules
		if contentRule == "" {
			rulesDomain, contentRule = getPredefinedScraperRules(entryUrl)
		}
		if contentRule != "" {
			content, ruleErr = ScrapContentUseRules(doc, rules)
			if ruleErr != nil {
				log.Printf(`get document by rule error rules:%s,domain:%s,%q`, rules, rulesDomain, ruleErr)
				return "", "", nil, "", "", "", "", publishedAtTimeStamp
			}
		}
	}
	if content != "" {
		content = rewrite.Rewriter(entryUrl, content, "add_dynamic_image")
		content = sanitizer.Sanitize(entryUrl, content)
	}

	if content != "" {
		article.Content = content
	}

	postContent := getPluginsPostContent(pluginsPath, entryUrl, article.Content)
	if postContent != "" {
		article.Content = postContent
	}

	pureContent := getPureContent(article.Content)

	return article.Content, pureContent, article.PublishedDate, article.Image, article.Title, author, article.Byline, publishedAtTimeStamp
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
