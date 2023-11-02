package dataCrawl

import (
	"net/http"

	"owl-webtoon/services/webtoon"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
)

// type dynamicCrawller struct {
// 	contextVar      context.Context
// 	cancelFunctions []context.CancelFunc
// }

type staticCrawller struct {
	document *goquery.Document
	//should close
	closeFunc func() error
}

type Crawller interface {
	getWebtoonInfos() ([]webtoon.Webtoon, error)
}

func newStaticCrawller(URL string) (*staticCrawller, error) {
	resp, err := http.Get(URL)
	if err != nil {
		return nil, errors.Wrap(err, "http get return error code")
	}

	html, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "invalid resp.Body")
	}
	return &staticCrawller{html, resp.Body.Close}, nil
}

// func newDynamicCrawller() (ret *dynamicCrawller) {
// 	ret = new(dynamicCrawller)
// 	opts := append(chromedp.DefaultExecAllocatorOptions[:],
// 		chromedp.DisableGPU,
// 		chromedp.Flag("headless", false), //headless를 false로 하면 브라우저가 뜨고, true로 하면 브라우저가 뜨지않는 headless 모드로 실행됨. 기본값은 true.
// 	)

// 	var cancelFunc context.CancelFunc

// 	ret.contextVar, cancelFunc = chromedp.NewExecAllocator(context.Background(), opts...)
// 	ret.cancelFunctions = append(ret.cancelFunctions, cancelFunc)

// 	ret.contextVar, cancelFunc = chromedp.NewContext(ret.contextVar)
// 	ret.cancelFunctions = append(ret.cancelFunctions, cancelFunc)
// 	return ret
// }
