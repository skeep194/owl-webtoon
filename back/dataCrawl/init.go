package dataCrawl

import (
	"context"

	"github.com/chromedp/chromedp"
)

var contextVar context.Context
var naverWebtoonURLPrefix string = "https://comic.naver.com/webtoon"
var cancelFunctions []context.CancelFunc

func init() {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.DisableGPU,
		chromedp.Flag("headless", false), //headless를 false로 하면 브라우저가 뜨고, true로 하면 브라우저가 뜨지않는 headless 모드로 실행됨. 기본값은 true.
	)

	var cancelFunc context.CancelFunc

	contextVar, cancelFunc = chromedp.NewExecAllocator(context.Background(), opts...)
	cancelFunctions = append(cancelFunctions, cancelFunc)

	contextVar, cancelFunc = chromedp.NewContext(contextVar)
	cancelFunctions = append(cancelFunctions, cancelFunc)
}
