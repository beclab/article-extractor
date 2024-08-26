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
	templateRawData := strings.NewReader(rawContent)
	doc, _ := goquery.NewDocumentFromReader(templateRawData)

	var content string
	var ruleErr error
	funcs := reflect.ValueOf(&templates.Template{})
	rulesDomain, contentRule := getPredefinedContentTemplateRules(entryUrl)
	if contentRule != "" {
		f := funcs.MethodByName(contentRule)
		res := f.Call([]reflect.Value{reflect.ValueOf(doc)})
		content = res[0].String()

	} else {
		contentRule = rules
		if contentRule == "" {
			rulesDomain, contentRule = getPredefinedScraperRules(entryUrl)
		}
		if contentRule != "" {
			content, ruleErr = templates.ScrapContentUseRules(doc, rules)
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

	postFuncs := reflect.ValueOf(&postExtractor.PostExtractorTemplate{})
	contentPreRule := getContentPostExtractorTemplateRules(entryUrl)
	if content != "" && contentPreRule != "" {
		f := postFuncs.MethodByName(contentPreRule)
		res := f.Call([]reflect.Value{reflect.ValueOf(content), reflect.ValueOf(feedUrl)})
		postContent := res[0].String()
		if postContent != "" {
			content = postContent
		}
	}
	pureContent := getPureContent(content)
	return content, pureContent
}

func RadioDetectionInArticle(rawContent, url string) string {
	templateRawData := strings.NewReader(rawContent)
	doc, _ := goquery.NewDocumentFromReader(templateRawData)
	funcs := reflect.ValueOf(&templates.Template{})
	_, mediaRule := getPredefinedMediaScraperRules(url)
	if mediaRule != "" {
		f := funcs.MethodByName(mediaRule)
		res := f.Call([]reflect.Value{reflect.ValueOf(url), reflect.ValueOf(doc)})
		//mediaContent := res[0].String()
		mediaUrl := res[1].String()
		mediaType := res[2].String()
		if mediaType == "audio" {
			return mediaUrl
		}
	}
	return ""
}
func ArticleReadabilityExtractor(rawContent, entryUrl, feedUrl, rules string, isrecommend bool) (string, string, *time.Time, string, string, int64, string, string, string) {
	var publishedAtTimeStamp int64 = 0
	templateRawData := strings.NewReader(rawContent)
	doc, _ := goquery.NewDocumentFromReader(templateRawData)

	rawData := strings.NewReader(rawContent)
	article, err := readability.FromReader(rawData, entryUrl)
	log.Printf("get readability article %s", entryUrl)
	if err != nil {
		log.Printf(`article extractor error %q`, err)
		return "", "", nil, "", "", publishedAtTimeStamp, "", "", ""
	}

	var content string
	var author string
	var ruleErr error
	var mediaContent string
	var mediaUrl string
	var mediaType string

	funcs := reflect.ValueOf(&templates.Template{})

	_, metadataRule := getPredefinedMetaDataTemplateRules(entryUrl)
	if metadataRule != "" {
		f := funcs.MethodByName(metadataRule)
		res := f.Call([]reflect.Value{reflect.ValueOf(doc)})
		author = res[0].String()
		/*published_at := res[1].String()
			if published_at != "" {
				ptime, parseErr := readability.ParseTime(published_at)
				if parseErr != nil {
					templateTime = &ptime
		}
			}*/
	} else if !strings.HasPrefix(feedUrl, "wechat") {
		author = templates.ScrapAuthorMetaData(doc)
	}

	_, mediaRule := getPredefinedMediaScraperRules(entryUrl)
	if mediaRule != "" {
		f := funcs.MethodByName(mediaRule)
		res := f.Call([]reflect.Value{reflect.ValueOf(entryUrl), reflect.ValueOf(doc)})
		mediaContent = res[0].String()
		mediaUrl = res[1].String()
		mediaType = res[2].String()
	}

	_, publishedAtRule := getPredefinedPublishedAtTimestampTemplateRules(entryUrl)
	log.Printf("current publishedAtRule %s", publishedAtRule)
	if publishedAtRule != "" {
		f := funcs.MethodByName(publishedAtRule)
		res := f.Call([]reflect.Value{reflect.ValueOf(doc)})
		publishedAtTimeStamp = res[0].Int()

	} else if !strings.HasPrefix(feedUrl, "wechat") {
		publishedAtTimeStamp = templates.ScrapPublishedAtTimeMetaData(doc)
	}
	if strings.HasPrefix(feedUrl, "wechat") {
		publishedAtTimeStamp = GetPublishedAtTimestampForWechat(rawContent, entryUrl)
	}

	rulesDomain, contentRule := getPredefinedContentTemplateRules(entryUrl)
	if contentRule != "" {
		f := funcs.MethodByName(contentRule)
		res := f.Call([]reflect.Value{reflect.ValueOf(doc)})
		content = res[0].String()

	} else {
		contentRule = rules
		if contentRule == "" {
			rulesDomain, contentRule = getPredefinedScraperRules(entryUrl)
		}
		if contentRule != "" {
			content, ruleErr = templates.ScrapContentUseRules(doc, rules)
			if ruleErr != nil {
				log.Printf(`get document by rule error rules:%s,domain:%s,%q`, rules, rulesDomain, err)
				return "", "", nil, "", "", publishedAtTimeStamp, "", "", ""
			}
		}
	}
	if content != "" {
		//readability.InsertToFile("before_add_dynamic_image.html", content)
		content = rewrite.Rewriter(entryUrl, content, "add_dynamic_image")

		content = sanitizer.Sanitize(entryUrl, content)
	}

	if content == "" {
		content = templates.GetArticleByDivClass(doc)
	}

	if content != "" || mediaType != "" {
		article.Content = content
	}

	postFuncs := reflect.ValueOf(&postExtractor.PostExtractorTemplate{})
	contentPreRule := getContentPostExtractorTemplateRules(entryUrl)
	if contentPreRule != "" {
		f := postFuncs.MethodByName(contentPreRule)
		res := f.Call([]reflect.Value{reflect.ValueOf(article.Content), reflect.ValueOf(feedUrl)})
		content := res[0].String()
		if content != "" {
			article.Content = content
		}
	}
	pureContent := getPureContent(article.Content)

	return article.Content, pureContent, article.PublishedDate, article.Image, author, publishedAtTimeStamp, mediaContent, mediaUrl, mediaType
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
