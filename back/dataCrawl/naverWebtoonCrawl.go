package dataCrawl

import (
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/pkg/errors"
)

type naverWebtoonCrawller struct {
	crawller
	baseURL string
}

func (naver naverWebtoonCrawller) GetWeekdayWebtoonURLs() (ret []string, err error) {
	var nodes []*cdp.Node

	err = chromedp.Run(naver.contextVar,
		chromedp.Navigate(naver.baseURL),
		chromedp.WaitVisible("li.DailyListItem__item--LP6_T"),
		chromedp.Nodes("li.DailyListItem__item--LP6_T > a.Poster__link--sopnC", &nodes),
	)
	if err != nil {
		return nil, errors.Wrap(err, "could not run chromedp")
	}

	ret = make([]string, len(nodes))
	for i, element := range nodes {
		ret[i] = element.Attributes[3]
	}
	return ret, nil
}

// func (naver naverWebtoonCrawller) getWebtoonInfoByURL(URL string) (ret *webtoon.Webtoon, err error) {

// }

// func (naver naverWebtoonCrawller) getWebtoonInfos() []webtoon.Webtoon {

// }
