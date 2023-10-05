package dataCrawl

import (
	"log"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

func GetWeekdayWebtoonURLs() []string {
	var nodes []*cdp.Node

	err := chromedp.Run(contextVar,
		chromedp.Navigate(naverWebtoonURLPrefix),
		chromedp.WaitVisible("li.DailyListItem__item--LP6_T"),
		chromedp.Nodes("li.DailyListItem__item--LP6_T > a.Poster__link--sopnC", &nodes),
	)
	if err != nil {
		log.Fatal(err)
	}

	ret := make([]string, len(nodes))
	for i, element := range nodes {
		ret[i] = element.Attributes[3]
	}
	return ret
}

func getWebtoonInfoByURL(URL string) {

}

func getWebtoonInfos() {
	
}