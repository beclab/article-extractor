package templates

import (
	"log"

	"github.com/PuerkitoBio/goquery"
)

func (t *Template) QuoraScrapContent(document *goquery.Document) string {
	contents := ""
	document.Find("div.spacing_log_question_page_ad").Each(func(i int, s *goquery.Selection) {
		RemoveNodes(s)
	})

	document.Find("div.puppeteer_test_answer_content").Each(func(i int, s *goquery.Selection) {
		parent := s.Parent()
		otherAnswer := parent.Find("div.spacing_log_originally_answered_banner").Length()
		relativeDivAnswer := parent.Find("div.puppeteer_test_question_title").Length()
		relativeSpanAnswer := parent.Find("span.puppeteer_test_question_title").Length()
		log.Printf("quora other answer otherAnswer: %d, relativeDivAnswer: %d, relativeSpanAnswer: %d", otherAnswer, relativeDivAnswer, relativeSpanAnswer)
		if otherAnswer == 0 && relativeDivAnswer == 0 && relativeSpanAnswer == 0 {
			author := ""
			prevNode := s.Prev()
			if prevNode.Length() > 0 {
				prevNode.Find("div > div > div > div > div > div > div > span > span >div >div >div >div >div > a >div > span >span").Each(func(i int, authors *goquery.Selection) {
					authorContent, e := authors.Html()
					if e == nil && author == "" {
						author = authorContent
						return
					}

				})
			}

			c1 := s.Children()
			if c1.Length() > 0 {
				cc1 := c1.Eq(0).Children()
				if cc1.Length() > 1 {
					ccc1 := cc1.Eq(1).Children()
					if ccc1.Length() == 1 {
						child := ccc1.Eq(0)
						content, err := child.Html()
						if err == nil {
							contents += "<div class='quora-item'><span class='q-author'><strong>" + author + "</strong></span>" + content + "<br></div>"
						}
					} else if ccc1.Length() > 1 {
						child := ccc1.Eq(1)
						content, err := child.Html()
						if err == nil {
							contents += "<div class='quora-item'><span class='q-author'><strong>" + author + "</strong></span>" + content + "<br></div>"
						}

					}
				}
			}
		}

	})
	return contents
}
