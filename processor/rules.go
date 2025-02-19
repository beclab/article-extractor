package processor

import (
	"fmt"
	"net/url"
	"strings"
)

// domain => CSS selectors
var contentPredefinedRules = map[string]string{
	"blog.cloudflare.com":  "div.post-content",
	"cbc.ca":               ".story-content",
	"darkreading.com":      "#article-main:not(header)",
	"developpez.com":       "div[itemprop=articleBody]",
	"dilbert.com":          "span.comic-title-name, img.img-comic",
	"explosm.net":          "div#comic",
	"financialsamurai.com": "article",
	"francetvinfo.fr":      ".text",
	"github.com":           "article.entry-content",
	"heise.de":             "header .article-content__lead, header .article-image, div.article-layout__content.article-content",
	"igen.fr":              "section.corps",
	"ikiwiki.iki.fi":       ".page.group",
	"ilpost.it":            ".entry-content",
	"ing.dk":               "section.body",
	"lapresse.ca":          ".amorce, .entry",
	"lemonde.fr":           "article",
	"lepoint.fr":           ".art-text",
	"lesjoiesducode.fr":    ".blog-post-content img",
	"lesnumeriques.com":    ".text",
	"linux.com":            "div.content, div[property]",
	"mac4ever.com":         "div[itemprop=articleBody]",
	"monwindows.com":       ".blog-post-body",
	"npr.org":              "#storytext",
	"oneindia.com":         ".io-article-body",
	"opensource.com":       "div[property]",
	"openingsource.org":    "article.suxing-popup-gallery",
	"osnews.com":           "div.newscontent1",
	"phoronix.com":         "div.content",
	"pseudo-sciences.org":  "#art_main",
	"quantamagazine.org":   ".outer--content, figure, script",
	"raywenderlich.com":    "article",
	"royalroad.com":        ".author-note-portlet,.chapter-content",
	"slate.fr":             ".field-items",
	"smbc-comics.com":      "div#cc-comicbody, div#aftercomic",
	"swordscomic.com":      "img#comic-image, div#info-frame.tab-content-area",
	"theoatmeal.com":       "div#comic",
	//"theregister.com":      "#top-col-story h2, #body",
	//"theverge.com":         "h2.inline:nth-child(2),h2.duet--article--dangerously-set-cms-markup,figure.w-full,div.duet--article--article-body-component",
	"turnoff.us":          "article.post-content",
	"universfreebox.com":  "#corps_corps",
	"version2.dk":         "section.body",
	"wdwnt.com":           "div.entry-content",
	"wired.com":           "div.grid-layout__content",
	"zeit.de":             ".summary, .article-body",
	"zdnet.com":           "div.storyBody",
	"pbfcomics":           "div#comic",
	"yahoo.com":           "div.caas-body",
	"kyivindependent.com": "div.c-content",
	//"news.mit.edu":        "div.news-article--content--body--inner",
}

var contentPostExtractorTemplateRules = map[string]string{
	"weixin.qq.com": "WechatPostExtractor",
	"espn.com":      "EspnPostExtractor",
}

var contentTemplatePredefinedRules = map[string]string{
	"apnews.com":                    "ApNewsScrapContent",
	"abcnews.go.com":                "AbcNewsScrapContent",
	"cnbc.com":                      "CNBCScrapContent",
	"bbc.co.uk":                     "BBCScrapContent",
	"bbc.com":                       "BBCScrapContent",
	"telegraph.co.uk":               "TelegraphScrapContent",
	"thestar.com":                   "TheStarScrapContent",
	"medium.com":                    "MediumScrapContent",
	"medium.datadriveninvestor.com": "MediumScrapContent",
	"cbsnews.com":                   "CbsNewsScrapContent",
	"news.sky.com":                  "SkyNewsScrapContent",
	"www.aljazeera.com":             "AljazeeraScrapContent",
	"themoscowtimes.com":            "MoscowTimesScrapContent",
	"themessenger.com":              "MessengerScrapContent",
	"euronews.com":                  "EuroNewsScrapContent",
	"huffpost.com":                  "HuffPostScrapContent",
	"dw.com":                        "DWScrapContent",
	"foxnews.com":                   "FoxNewsScrapContent",
	"pravda.com":                    "PravdaScrapContent",
	"time.com":                      "TimeScrapContent",
	"theguardian.com":               "TheguardianScrapContent",
	"reuters.com":                   "ReutersScrapContent",
	"abc.net.au":                    "AbcNetAUScrapContent",
	"yahoo.com":                     "YahoocrapContent",
	"nbcnews.com":                   "NbcNewsScrapContent",
	"cncf.io":                       "CNCFScrapContent",
	"deepmind.google":               "DeepMindScrapContent",
	"digitaltrends.com":             "DigitalTrendsScrapContent",
	"nypost.com":                    "NYpostScrapContent",
	"techcrunch.com":                "TechCrunchScrapContent",
	"theverge.com":                  "ThevergeScrapContent",

	"theregister.com":  "TheRegisterScrapContent",
	"dazeddigital.com": "DazeddigitalScrapContent",
	"deadline.com":     "DeadlineScrapContent",
	//"eonline.com":                   "EOnlineScrapContent",
	"filmschoolrejects.com":        "FilmSchoolRejectsScrapContent",
	"independent.co.uk":            "IndependentUKScrapContent",
	"skysports.com":                "SkySportsScrapContent",
	"sbnation.com":                 "SbnationScrapContent",
	"cbssports.com":                "CBSsportsScrapContent",
	"scmp.com":                     "SCMPScrapContent",
	"cnn.com":                      "CNNScrapContent",
	"businesslive.co.za":           "BusinessLiveScrapContent",
	"smallbiztrends.com":           "SmallBizTrendsScrapContent",
	"hbr.org":                      "HBRScrapContent",
	"entrepreneur.com":             "EntrepreneurScrapContent",
	"businessinsider.com":          "BusinessInsiderScrapContent",
	"mattturck.com":                "MattturckScrapContent",
	"cfainstitute.org":             "CFainstituteScrapContent",
	"hunterwalk.com":               "HunterWalkScrapContent",
	"thisisgoingtobebig.com":       "ThisisGoingtobeBIGScrapContent",
	"ign.com":                      "IGNScrapContent",
	"screenrant.com":               "ScreenrantScrapContent",
	"vice.com":                     "ViceScrapContent",
	"variety.com":                  "VarietyScrapContent",
	"avclub.com":                   "AVClubScrapContent",
	"stereogum.com":                "StereogumScrapContent",
	"pitchfork.com":                "PitchForkScrapContent",
	"geektyrant.com":               "GeektyrantScrapContent",
	"advanced-television.com":      "AdTelevisionScrapContent",
	"bleacherreport.com":           "BleadherReportScrapContent",
	"foxsports.com":                "FoxSportsScrapContent",
	"polygon.com":                  "PolygonScrapContent",
	"kotaku.com":                   "KotakuScrapContent",
	"tvline.com":                   "TVLineScrapContent",
	"npr.org":                      "NprScrapContent",
	"microsoft.com":                "MicrosoftScrapContent",
	"benzinga.com":                 "BenzingaScrapContent",
	"vox.com":                      "VoxScrapContent",
	"financialpost.com":            "FinancialPostScrapContent",
	"ibtimes.co.uk":                "IbtimesScrapContent",
	"economictimes.indiatimes.com": "IndiatimesScrapContent",
	"ycombinator.com":              "YcombinatorScrapContent",
	"visualcapitalist.com":         "VisualcapitalistScrapContent",
	"fastcompany.com":              "FastcompanyScrapContent",
	"futurism.com":                 "FuturismScrapContent",
	"gizmodo.com":                  "GizmodoScrapContent",
	"interestingengineering.com":   "InterestingengineeringScrapContent",
	"mirror.co.uk":                 "MirrorScrapContent",
	"newatlas.com":                 "NewatlasScrapContent",
	"techradar.com":                "TechradarScrapContent",
	"techspot.com":                 "TechspotScrapContent",
	"news.mit.edu":                 "MITScrapContent",
	"pagesix.com":                  "PagesixScrapContent",
	"thedailybeast.com":            "ThedailybeastScrapContent",
	"radiotimes.com":               "RadiotimesScrapContent",
	"hollywoodreporter.com":        "HollywoodreporterScrapContent",
	"usmagazine.com":               "UsmagazineScrapContent",
	"thewrap.com":                  "ThewrapScrapContent",
	"the-sun.com":                  "TheSunScrapContent",
	"thesun.co.uk":                 "TheSunScrapContent",
	"themirror.com":                "TheMirrorScrapContent",
	"dailymail.co.uk":              "DailymailScrapContent",
	"ew.com":                       "EWScrapContent",
	"storyfm.cn":                   "StoryFMScrapContent",
	"bilibili.com":                 "BilibiliScrapContent",
	"podbean.com":                  "PodBeanScrapContent",
	"xiaoyuzhoufm.com":             "XiaoyuzhouFMScrapContent",
	"youtube.com":                  "YoutubeScrapContent",
	"vimeo.com":                    "VimeoScrapContent",
	"rumble.com":                   "RumbleScrapContent",
	"spreaker.com":                 "SpreakerScrapContent",
	"pinterest.com":                "PinterestScrapContent",
	"acfun.cn":                     "ACFunScrapContent",
	"wsj.com":                      "WsjScrapContent",
	"ft.com":                       "FTScrapContent",
	"notion.site":                  "NotionScrapContent",
	"zhihu.com":                    "ZhihuScrapContent",
	"wolai.com":                    "WolaiScrapContent",
	//"screencrush.com":              "ScreencrushScrapContent",
	/*"espn.com": "EspnScrapContent",
	"nbcsports.com":      "NBCSportsScrapContent",
	"deadspin.com":       "DeadspinScrapContent",
	"skynews.com":              "SkyNewsScrapContent",
	*/
}

var mediaTemplatePredefinedRules = map[string]string{
	"storyfm.cn":       "StoryFMMediaContent",
	"bilibili.com":     "BilibiliMediaContent",
	"podbean.com":      "PodBeanMediaContent",
	"xiaoyuzhoufm.com": "XiaoyuzhouFMMediaContent",
	"youtube.com":      "YoutubeMediaContent",
	"vimeo.com":        "VimeoMediaContent",
	"rumble.com":       "RumbleMediaContent",
	"spreaker.com":     "SpreakerMediaContent",
	"pinterest.com":    "PinterestMediaContent",
	"acfun.cn":         "ACFunMediaContent",
}

var metadataTemplatePredefinedRules = map[string]string{
	"eonline.com":         "EonlineScrapMetaData",
	"slashfilm.com":       "SlashfilmScrapMetaData",
	"abcnews.go.com":      "AbcNewsScrapMetaData",
	"apnews.com":          "ApnNewsScrapMetaData",
	"www.aljazeera.com":   "AljazeeraScrapMetaData",
	"news.sky.com":        "SkyNewsScrapMetaData",
	"yahoo.com":           "YahooNewsScrapMetaData",
	"abc.net.au":          "AbcNetAUScrapMetaData",
	"cbsnews.com":         "CbsNewsScrapMetaData",
	"cnbc.com":            "CnbcScrapMetaData",
	"dw.com":              "DWScrapMetaData",
	"euronews.com":        "EuroNewsScrapMetaData",
	"foxnews.com":         "FoxNewsScrapMetaData",
	"huffpost.com":        "HuffPostScrapMetaData",
	"nbcnews.com":         "NbcNewsScrapMetaData",
	"ndtv.com":            "NdtvNewsScrapMetaData",
	"pravda.com":          "PravdaScrapMetaData",
	"themoscowtimes.com":  "ThemoscowtimesScrapMetaData",
	"themessenger.com":    "ThemessengerScrapMetaData",
	"theguardian.com":     "TheguardianScrapMetaData",
	"www.bbc.co.uk/news":  "BBCNewsScrapMetaData",
	"npr.org":             "NprScrapMetaData",
	"stereogum.com":       "StereogumScrapMetaData",
	"www.vice.com":        "ViceScrapMetaData",
	"a16z.com":            "A16ZScrapMetaData",
	"a16zcrypto.com":      "A16ZCrptoScrapMetaData",
	"businessinsider.com": "BusinessinsiderScrapMetaData",
	"foxbusiness.com":     "FoxbusinessScrapMetaData",
	"businesslive.co.za":  "BusinessliveScrapMetaData",
	"edition.cnn.com":     "EditionCnnScrapMetaData",
	"money.cnn.com":       "EditionCnnScrapMetaData",
	"skysports.com":       "SkySportsScrapMetaData",
	"www.bbc.com/sport":   "BBCSportsScrapMetaData",
	"www.bbc.co.uk/sport": "BBCSportsScrapMetaData",
	"cbssports.com":       "CBSSportsScrapMetaData",
	".espn.com":           "ESPNScrapMetaData",
	"foxsports.com":       "FoxsportsScrapMetaData",
	"hbr.org":             "HBRScrapMetaData",
	"nbcsports.com":       "NBCSPortScrapMetaData",
	"cncf.io":             "CNCFScrapMetaData",
	"time.com":            "TimeScrapMetaData",
	"deepmind.google":     "DeepMindScrapMetaData",
	"screenrant.com":      "ScreenrantScrapMetaData",
	"deadline.com":        "DeadlineScrapMetaData",
	"variety.com":         "VarietyScrapMetaData",
	"newatlas.com":        "NewatlasScrapMetaData",
	"koreatimes.co.kr":    "KoreatimesScrapMetaData",
	"pinterest.com":       "PinterestScrapMetaData",
	"acfun.cn":            "ACFunScrapMetaData",
	"zhihu.com":           "ZhihuScrapMetaData",
}

var publishedAtTimeStampTemplatePredefinedRules = map[string]string{
	"slashfilm.com":              "SlashfilmNewsPublishedAtTimeFromScriptMetadata",
	"abcnews.go.com":             "CommonGetPublishedAtTimestampSingleJson",
	"apnews.com":                 "ApNewsCommonGetPublishedAtTimestamp",
	"www.aljazeera.com":          "CommonGetPublishedAtTimestampSingleJson",
	"news.sky.com":               "SkyNewsPublishedAtTimeFromScriptMetadata",
	"yahoo.com":                  "CommonGetPublishedAtTimestampSingleJson",
	"abc.net.au":                 "CommonGetPublishedAtTimestampSingleJson",
	"cbsnews.com":                "CbsnewsWorldGetPublishedAtTimestampSingleJson",
	"cnbc.com":                   "CnbcPublishedAtTimeFromScriptMetadata",
	"dw.com":                     "CommonGetPublishedAtTimestampSingleJson",
	"euronews.com":               "EuroNewsGetPublishedAtTimeStampStruct",
	"foxnews.com":                "CommonGetPublishedAtTimestampSingleJson",
	"huffpost.com":               "CommonGetPublishedAtTimestampSingleJson",
	"nbcnews.com":                "CommonGetPublishedAtTimestampSingleJson",
	"ndtv.com":                   "NdtvGetPublishedAtTimestamp",
	"pravda.com":                 "CommonGetPublishedAtTimestampSingleJson",
	"themoscowtimes.com":         "CommonGetPublishedAtTimestampSingleJson",
	"themessenger.com":           "TheMessengerGetPublishedAtTimestampSingleJson",
	"theguardian.com":            "CommonGetPublishedAtTimestampMultipleJson",
	"www.bbc.co.uk/news":         "BBCNewsPublishedAtTimeFromScriptMetadata",
	"time.com":                   "TimePublishedAtTimeFromScriptMetadata",
	"eonline.com":                "EonlinePublishedAtTimeFromScriptMetadata",
	"npr.org":                    "NprPublishedAtTimeFromScriptMetadata",
	"stereogum.com":              "StereogumPublishedAtTimeFromScriptMetadata",
	"www.vice.com":               "VicePublishedAtTimeFromScriptMetadata",
	"a16z.com":                   "A16ZPublishedAtTimeFromScriptMetadata",
	"businessinsider.com":        "BusinessinsiderPublishedAtTimeFromScriptMetadata",
	"foxbusiness.com":            "FoxbusinessPublishedAtTimeFromScriptMetadata",
	"businesslive.co.za":         "BusinesslivePublishedAtTimeFromScriptMetadata",
	"edition.cnn.com":            "EditionCnnPublishedAtTimeFromScriptMetadata",
	"money.cnn.com":              "EditionCnnPublishedAtTimeFromScriptMetadata",
	"skysports.com":              "SkySportsPublishedAtTimeFromScriptMetadata",
	"www.bbc.com/sport":          "BBCSportsPublishedAtTimeFromScriptMetadata",
	"www.bbc.co.uk/sport":        "BBCSportsPublishedAtTimeFromScriptMetadata",
	"cbssports.com":              "CBSSportPublishedAtTimeFromScriptMetadata",
	"espn.com":                   "ESPNPublishedAtTimeFromScriptMetadata",
	"foxsports.com":              "FoxsportsPublishedAtTimeFromScriptMetadata",
	"hbr.org":                    "HBRPublishedAtTimeFromScriptMetadata",
	"nbcsports.com":              "NBCSportsPublishedAtTimeFromScriptMetadata",
	"cncf.io":                    "CNCFPublishedAtTimeFromScriptMetadata",
	"deepmind.google":            "DeepMindPublishedAtTimeFromScriptMetadata",
	"screenrant.com":             "ScreenrantPublishedAtTimeFromScriptMetadata",
	"deadline.com":               "DeadlinePublishedAtTimeFromScriptMetadata",
	"variety.com":                "VarietyPublishedAtTimeFromScriptMetadata",
	"interestingengineering.com": "InterestingengineeringPublishedAtTimeFromScriptMetadata",
	"acfun.cn":                   "ACFunPublishedAtTimeFromScriptMetadata",
}

func getPredefinedPublishedAtTimestampTemplateRules(websiteURL string) (string, string) {
	urlDomain := domain(websiteURL)
	for domain, rules := range publishedAtTimeStampTemplatePredefinedRules {
		if strings.Contains(websiteURL, domain) {
			return domain, rules
		}
	}
	for domain, rules := range publishedAtTimeStampTemplatePredefinedRules {
		if strings.Contains(urlDomain, domain) {
			return domain, rules
		}
	}
	return "", ""
}

func getContentPostExtractorTemplateRules(websiteURL string) string {
	urlDomain := domain(websiteURL)
	for url, rules := range contentPostExtractorTemplateRules {
		if strings.Contains(urlDomain, url) {
			return rules
		}
	}
	return ""
}

func getPredefinedMediaScraperRules(websiteURL string) (string, string) {
	urlDomain := domain(websiteURL)

	for domain, rules := range mediaTemplatePredefinedRules {
		if strings.Contains(urlDomain, domain) {
			return domain, rules
		}
	}
	return "", ""
}

func getPredefinedContentTemplateRules(websiteURL string) (string, string) {
	urlDomain := domain(websiteURL)
	for domain, rules := range contentTemplatePredefinedRules {
		if strings.Contains(urlDomain, domain) {
			return domain, rules
		}
	}
	return "", ""
}

func getPredefinedMetaDataTemplateRules(websiteURL string) (string, string) {
	urlDomain := domain(websiteURL)
	fmt.Printf("+++++++++++++++++++++ websiteURL %s\n", websiteURL)
	for domain, rules := range metadataTemplatePredefinedRules {
		if strings.Contains(websiteURL, domain) {
			return domain, rules
		}
	}
	for domain, rules := range metadataTemplatePredefinedRules {
		if strings.Contains(urlDomain, domain) {
			return domain, rules
		}
	}
	return "", ""
}

func getPredefinedScraperRules(websiteURL string) (string, string) {
	urlDomain := domain(websiteURL)

	for domain, rules := range contentPredefinedRules {
		if strings.Contains(urlDomain, domain) {
			return domain, rules
		}
	}
	return "", ""
}

func domain(websiteURL string) string {
	parsedURL, err := url.Parse(websiteURL)
	if err != nil {
		return websiteURL
	}

	return parsedURL.Host
}
