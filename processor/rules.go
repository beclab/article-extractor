package processor

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var contentPostExtractorTemplateRules = map[string]string{
	"weixin.qq.com": "WechatPostExtractor",
	"espn.com":      "EspnPostExtractor",
}

var contentTemplatePredefinedRules = map[string]string{
	"a16zcrypto.com":          "A16ZCrptoExtractorMetaInfo",
	"abc.net.au":              "AbcNetExtractorMetaInfo",
	"abcnews.go.com":          "AbcNewsExtractorMetaInfo",
	"acfun.cn":                "ACFunExtractorMetaInfo",
	"advanced-television.com": "AdTelevisionExtractorMetaInfo",
	"aljazeera.com":           "AljazeeraExtractorMetaInfo",
	"apnews.com":              "ApNewsExtractorMetaInfo",
	"avclub.com":              "AVClubExtractorMetaInfo",
	"bbc.co.uk":               "BBCExtractorMetaInfo",
	"bbc.com":                 "BBCExtractorMetaInfo",
	"benzinga.com":            "BenzingaExtractorMetaInfo",
	"bilibili.com":            "BilibiliExtractorMetaInfo",
	"bleacherreport.com":      "BleadherReportExtractorMetaInfo",
	"businessinsider.com":     "BusinessInsiderExtractorMetaInfo",
	"businesslive.co.za":      "BusinessLiveExtractorMetaInfo",
	"cbsnews.com":             "CbsNewsExtractorMetaInfo",
	"cbssports.com":           "CBSSportExtractorMetaInfo",
	"cfainstitute.org":        "CFainstitutExtractorMetaInfo",
	"cnbc.com":                "CNBCExtractorMetaInfo",
	"cncf.io":                 "CNCFExtractorMetaInfo",
	"cnn.com":                 "CNNExtractorMetaInfo",
	"dailymail.co.uk":         "DailymailExtractorMetaInfo",
	"dazeddigital.com":        "DazeddigitalExtractorMetaInfo",
	"deadline.com":            "DeadlineExtractorMetaInfo",
	"deadspin.com":            "DeadspinExtractorMetaInfo",
	"deepmind.google":         "DeepMindExtractorMetaInfo",
	"digitaltrends.com":       "DigitalTrendsExtractorMetaInfo",
	"douban.com":              "DoubanExtractorMetaInfo",
	"dw.com":                  "DWExtractorMetaInfo",
	"entrepreneur.com":        "EntrepreneurExtractorMetaInfo",
	"eonline.com":             "EOnlineExtractorMetaInfo",
	"espn.com":                "EspnExtractorMetaInfo",
	"euronews.com":            "EuroNewsExtractorMetaInfo",
	"ew.com":                  "EWExtractorMetaInfo",
	"fandom.com":              "FandomExtractorMetaInfo",
	"fastcompany.com":         "FastcompanyExtractorMetaInfo",
	"feishu.cn":               "FeishuExtractorMetaInfo",
	"filmschoolrejects.com":   "FilmSchoolRejectsExtractorMetaInfo",
	"financialpost.com":       "FinancialPostExtractorMetaInfo",
	"foxnews.com":             "FoxNewsExtractorMetaInfo",
	"foxsports.com":           "FoxSportsExtractorMetaInfo",
	"ft.com":                  "FTExtractorMetaInfo",
	"futurism.com":            "FuturismExtractorMetaInfo",
	"geektyrant.com":          "GeektyrantExtractorMetaInfo",
	//ghost
	"gizmodo.com":                   "GizmodoExtractorMetaInfo",
	"hbr.org":                       "HBRExtractorMetaInfo",
	"hollywoodreporter.com":         "HollywoodreporterExtractorMetaInfo",
	"huffpost.com":                  "HuffPostExtractorMetaInfo",
	"hunterwalk.com":                "HunterWalkExtractorMetaInfo",
	"ibtimes.co.uk":                 "IbtimesExtractorMetaInfo",
	"ign.com":                       "IGNExtractorMetaInfo",
	"independent.co.uk":             "IndependentUKExtractorMetaInfo",
	"economictimes.indiatimes.com":  "IndiatimesExtractorMetaInfo",
	"interestingengineering.com":    "InterestingengineeringExtractorMetaInfo",
	"jianshu.com":                   "JianshuExtractorMetaInfo",
	"koreatimes.co.kr":              "KoreatimesExtractorMetaInfo",
	"kotaku.com":                    "KotakuExtractorMetaInfo",
	"lizhi.fm":                      "LizhiExtractorMetaInfo",
	"mattturck.com":                 "MattturckExtractorMetaInfo",
	"medium.com":                    "MediumExtractorMetaInfo",
	"medium.datadriveninvestor.com": "MediumExtractorMetaInfo",
	"themessenger.com":              "MessengerExtractorMetaInfo",
	"microsoft.com":                 "MicrosoftExtractorMetaInfo",
	"mirror.co.uk":                  "MirrorExtractorMetaInfo",
	"news.mit.edu":                  "MITExtractorMetaInfo",
	"themoscowtimes.com":            "MoscowTimesExtractorMetaInfo",
	"nbcnews.com":                   "NbcNewsExtractorMetaInfo",
	"nbcsports.com":                 "NBCSportsExtractorMetaInfo",
	"newatlas.com":                  "NewatlasExtractorMetaInfo",
	"notion.site":                   "NotionExtractorMetaInfo",
	"npr.org":                       "NprExtractorMetaInfo",
	"nypost.com":                    "NYpostExtractorMetaInfo",

	"okjike.com":             "OKjikeExtractorMetaInfo",
	"pagesix.com":            "PagesixExtractorMetaInfo",
	"pinterest.com":          "PinterestExtractorMetaInfo",
	"pitchfork.com":          "PitchForkExtractorMetaInfo",
	"podbean.com":            "PodBeanExtractorMetaInfo",
	"polygon.com":            "PolygonExtractorMetaInfo",
	"pravda.com":             "PravdaExtractorMetaInfo",
	"quora.com":              "QuoraExtractorMetaInfo",
	"radiotimes.com":         "RadiotimesExtractorMetaInfo",
	"reddit.com":             "RedditExtractorMetaInfo",
	"reuters.com":            "ReutersExtractorMetaInfo",
	"rumble.com":             "RumbleExtractorMetaInfo",
	"sbnation.com":           "SbnationExtractorMetaInfo",
	"scmp.com":               "SCMPExtractorMetaInfo",
	"screencrush.com":        "ScreencrushExtractorMetaInfo",
	"screenrant.com":         "ScreenrantExtractorMetaInfo",
	"news.sky.com":           "SkyNewsExtractorMetaInfo",
	"skynews.com":            "SkyNewsExtractorMetaInfo",
	"skysports.com":          "SkySportsExtractorMetaInfo",
	"smallbiztrends.com":     "SmallBizTrendsExtractorMetaInfo",
	"spreaker.com":           "SpreakerExtractorMetaInfo",
	"stereogum.com":          "StereogumExtractorMetaInfo",
	"storyfm.cn":             "StoryFMExtractorMetaInfo",
	"techcrunch.com":         "TechCrunchExtractorMetaInfo",
	"techradar.com":          "TechradarExtractorMetaInfo",
	"techspot.com":           "TechspotExtractorMetaInfo",
	"telegraph.co.uk":        "TelegraphExtractorMetaInfo",
	"thedailybeast.com":      "ThedailybeastExtractorMetaInfo",
	"theguardian.com":        "TheguardianExtractorMetaInfo",
	"themirror.com":          "TheMirrorExtractorMetaInfo",
	"thestar.com":            "TheStarExtractorMetaInfo",
	"theregister.com":        "TheRegisterExtractorMetaInfo",
	"the-sun.com":            "TheSunExtractorMetaInfo",
	"thesun.co.uk":           "TheSunExtractorMetaInfo",
	"theverge.com":           "ThevergeExtractorMetaInfo",
	"thewrap.com":            "ThewrapExtractorMetaInfo",
	"thisisgoingtobebig.com": "ThisisGoingtobeBigExtractorMetaInfo",
	"time.com":               "TimeExtractorMetaInfo",
	"tvline.com":             "TVLineExtractorMetaInfo",
	"usmagazine.com":         "UsmagazineExtractorMetaInfo",
	"v2ex.com":               "V2exExtractorMetaInfo",
	"variety.com":            "VarietyExtractorMetaInfo",
	"vice.com":               "ViceExtractorMetaInfo",
	"vimeo.com":              "VimeoExtractorMetaInfo",
	"visualcapitalist.com":   "VisualcapitalistExtractorMetaInfo",
	"vox.com":                "VoxExtractorMetaInfo",
	"wolai.com":              "WolaiExtractorMetaInfo",
	"wsj.com":                "WsjExtractorMetaInfo",
	"xhslink.com":            "XhsExtractorMetaInfo",
	"xiaoyuzhoufm.com":       "XiaoyuzhouFMExtractorMetaInfo",
	"ximalaya.com":           "XimalayaExtractorMetaInfo",
	"yahoo.com":              "YahooExtractorMetaInfo",
	"ycombinator.com":        "YcombinatorExtractorMetaInfo",
	"youtube.com":            "YoutubeExtractorMetaInfo",
	"zhihu.com":              "ZhihuExtractorMetaInfo",
}

var DownloadTypeUrlTemplatedRules = map[string]string{
	"manybooks.net":      "ManyBooksDownloadType", //need cookies
	"standardebooks.org": "StandardebooksDownloadType",
	"z-library.gs":       "ZLibraryDownloadType", //need cookies
}

func getPredefinedRules(websiteURL string, doc *goquery.Document) string {
	urlDomain := domain(websiteURL)
	isSubstack := false
	if strings.Contains(urlDomain, "substack.com") || doc.Find(`link[rel="preconnect"][href="https://substackcdn.com"]`).Length() > 0 {
		isSubstack = true
	}
	if isSubstack {
		return "SubStackExtractorMetaInfo"
	}
	for domain, rules := range contentTemplatePredefinedRules {
		if strings.Contains(urlDomain, domain) {
			return rules
		}
	}
	return ""
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

func getDownloadTypeByUrlRules(websiteURL string) (string, string) {
	urlDomain := domain(websiteURL)
	for domain, rules := range DownloadTypeUrlTemplatedRules {
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
