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

func getRulesFromContent(websiteURL string, doc *goquery.Document) string {
	urlDomain := domain(websiteURL)
	isSubstack := false
	if strings.Contains(urlDomain, "substack.com") || doc.Find(`link[rel="preconnect"][href="https://substackcdn.com"]`).Length() > 0 {
		isSubstack = true
	}
	if isSubstack {
		return "SubStackScrapContent"
	}
	/*isGhost := false
	content, exists := doc.Find(`meta[name="generator"]`).Attr("content")
	if exists {
		if strings.HasPrefix(content, "Ghost") && doc.Find("section.gh-content").Length() > 0 {
			isGhost = true
		}
	}

	if isGhost {
		return "GhostScrapContent"
	}*/
	return ""
}
func ArticleContentExtractor(rawContent, entryUrl, feedUrl, rules string) (string, string) {
	templateRawData := strings.NewReader(rawContent)
	doc, _ := goquery.NewDocumentFromReader(templateRawData)

	var content string
	var ruleErr error
	funcs := reflect.ValueOf(&templates.Template{})
	rulesDomain := ""
	contentRule := getRulesFromContent(entryUrl, doc)
	if contentRule == "" {
		rulesDomain, contentRule = getPredefinedContentTemplateRules(entryUrl)
	}
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

func ExceptYTdlpDownloadQueryInArticle(rawContent, url string) (string, string) {
	templateRawData := strings.NewReader(rawContent)
	doc, _ := goquery.NewDocumentFromReader(templateRawData)
	funcs := reflect.ValueOf(&templates.Template{})
	_, mediaRule := getPredefinedMediaScraperRules(url)
	if mediaRule != "" {
		f := funcs.MethodByName(mediaRule)
		res := f.Call([]reflect.Value{reflect.ValueOf(url), reflect.ValueOf(doc)})
		//mediaContent := res[0].String()
		downloadUrl := res[1].String()
		downloadType := res[2].String()
		if downloadType == "audio" || downloadType == "ebook" || downloadType == "pdf" {
			return downloadUrl, downloadType
		}
	}
	return "", ""
}
func ArticleReadabilityExtractor(rawContent, entryUrl, feedUrl, rules string, isrecommend bool) (string, string, *time.Time, string, string, string, int64, string, string, string) {
	var publishedAtTimeStamp int64 = 0
	templateRawData := strings.NewReader(rawContent)
	doc, _ := goquery.NewDocumentFromReader(templateRawData)

	entryDomain := domain(entryUrl)
	rawData := strings.NewReader(rawContent)
	article, err := readability.FromReader(rawData, entryUrl)
	log.Printf("get readability article %s", entryUrl)
	if err != nil {
		log.Printf(`article extractor error %q`, err)
		return "", "", nil, "", "", "", publishedAtTimeStamp, "", "", ""
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
	} else { //if !strings.Contains(entryUrl, "weixin.qq.com") { //else if !strings.HasPrefix(feedUrl, "wechat") {
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

	} else if !strings.Contains(entryUrl, "weixin.qq.com") { //else if !strings.HasPrefix(feedUrl, "wechat") {
		publishedAtTimeStamp = templates.ScrapPublishedAtTimeMetaData(doc)
	}
	if strings.Contains(entryUrl, "weixin.qq.com") { //if strings.HasPrefix(feedUrl, "wechat") {
		publishedAtTimeStamp = GetPublishedAtTimestampForWechat(rawContent, entryUrl)
		if author == "" {
			author = GetAuthorForWechat(doc)
		}
	}

	rulesDomain := ""
	contentRule := getRulesFromContent(entryUrl, doc)
	if contentRule == "" {
		rulesDomain, contentRule = getPredefinedContentTemplateRules(entryUrl)
	}
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
				return "", "", nil, "", "", "", publishedAtTimeStamp, "", "", ""
			}
		}
	}
	if content != "" {
		//readability.InsertToFile("before_add_dynamic_image.html", content)
		if !strings.Contains(entryDomain, "douban.com") {
			content = rewrite.Rewriter(entryUrl, content, "add_dynamic_image")
		}
		if strings.Contains(entryDomain, "okjike.com") || strings.Contains(entryDomain, "vimeo.com") ||
			strings.Contains(entryDomain, "fandom.com") ||
			strings.Contains(entryDomain, "notion.site") || strings.Contains(entryDomain, "quora.com") {
			//不进行santitize
			if strings.Contains(entryDomain, "notion.site") {
				article.Title = doc.Find("title").Text()
			}
		} else {
			content = sanitizer.Sanitize(entryUrl, content)
		}
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
	/*checkArticleImage := getArticleImage(rawContent, entryUrl)
	if checkArticleImage != "" {
		article.Image = checkArticleImage
	}*/

	if strings.Contains(entryDomain, "reddit.com") {
		doc.Find("shreddit-post").Each(func(i int, s *goquery.Selection) {
			if title, exists := s.Attr("post-title"); exists {
				article.Title = strings.TrimSpace(title)
			}
		})
	}
	return article.Content, pureContent, article.PublishedDate, article.Image, article.Title, author, publishedAtTimeStamp, mediaContent, mediaUrl, mediaType
}

/*
	func getArticleImage(content, url string) string {
		articleImage := ""
		if strings.Contains(url, "bilibili.com/bangumi/play") {
			ep := ""
			re := regexp.MustCompile(`play/ep(\d+)`)
			match := re.FindStringSubmatch(url)
			if len(match) > 1 {
				ep = match[1]
			}
			log.Printf("get bili bangumi image:%s", ep)
			if ep != "" {
				doc, _ := goquery.NewDocumentFromReader(strings.NewReader(content))
				var session interface{}
				doc.Find("script[type='application/json']").Each(func(i int, s *goquery.Selection) {
					scriptContent := strings.TrimSpace(s.Text())
					var metaData map[string]interface{}
					unmarshalErr := json.Unmarshal([]byte(scriptContent), &metaData)
					if unmarshalErr != nil {
						log.Printf("convert  unmarshalError %v", unmarshalErr)
						return
					}
					if props, ok := metaData["props"]; ok {
						if pageProps, ok := props.(map[string]interface{})["pageProps"]; ok {
							if dehydratedState, ok := pageProps.(map[string]interface{})["dehydratedState"]; ok {
								if queries, ok := dehydratedState.(map[string]interface{})["queries"]; ok {
									queriesArr := queries.([]interface{})
									if len(queriesArr) > 1 {
										if state, ok := queriesArr[1].(map[string]interface{})["state"]; ok {
											if stateData, ok := state.(map[string]interface{})["data"]; ok {
												session = stateData.(map[string]interface{})["seasons"]
											}
										}
									}
								}
							}
						}
					}
				})
				if session != nil {
					for _, seasonDetail := range session.([]interface{}) {
						sessionMap := seasonDetail.(map[string]interface{})
						if new_ep, ok := sessionMap["new_ep"]; ok {
							new_ep_data := new_ep.(map[string]interface{})
							id, idok := new_ep_data["id"]
							cover, coverok := new_ep_data["cover"]
							idstr := fmt.Sprintf("%.0f", id.(float64))
							if idok && coverok && idstr == ep {
								articleImage = cover.(string)
								return articleImage
							}
						}
					}
				}
			}
		}
		return articleImage
	}
*/
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
	author := ""
	document.Find("div#meta_content>span.rich_media_meta_text").Each(func(i int, s *goquery.Selection) {
		content := s.Text()
		content = strings.TrimSpace(content)
		content = strings.ReplaceAll(content, "\n", "")
		author = content
	})
	if author == "" {
		document.Find("div#meta_content>span.rich_media_meta_nickname").Each(func(i int, s *goquery.Selection) {
			content := s.Text()
			content = strings.TrimSpace(content)
			content = strings.ReplaceAll(content, "\n", "")
			author = content
		})
	}

	return author

}
