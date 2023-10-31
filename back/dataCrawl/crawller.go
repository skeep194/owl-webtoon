package dataCrawl

import (
	"context"

	"owl-webtoon/services/webtoon"

	"github.com/chromedp/chromedp"
)

type crawller struct {
	contextVar      context.Context
	cancelFunctions []context.CancelFunc
}

type Crawller interface {
	getWebtoonInfos() []webtoon.Webtoon
}

// var contextVar context.Context
// var naverWebtoonURLPrefix string = "https://comic.naver.com/webtoon"
// var cancelFunctions []context.CancelFunc

func newCrawller() (ret *crawller) {
	ret = new(crawller)
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.DisableGPU,
		chromedp.Flag("headless", false), //headless를 false로 하면 브라우저가 뜨고, true로 하면 브라우저가 뜨지않는 headless 모드로 실행됨. 기본값은 true.
	)

	var cancelFunc context.CancelFunc

	ret.contextVar, cancelFunc = chromedp.NewExecAllocator(context.Background(), opts...)
	ret.cancelFunctions = append(ret.cancelFunctions, cancelFunc)

	ret.contextVar, cancelFunc = chromedp.NewContext(ret.contextVar)
	ret.cancelFunctions = append(ret.cancelFunctions, cancelFunc)
	return ret
}
