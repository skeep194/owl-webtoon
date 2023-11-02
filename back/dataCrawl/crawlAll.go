package dataCrawl

import (
	"owl-webtoon/services/webtoonService"

	"github.com/pkg/errors"
)

var crawllers = []Crawller{
	naverWebtoonCrawller{},
}

func Crawl() (err error) {
	for _, crawller := range crawllers {
		var webtoons []webtoonService.Webtoon
		webtoons, err = crawller.getWebtoonInfos()
		if err != nil {
			err = errors.Wrap(err, "some error from get webtoon infos")
		}
		webtoonService.InsertWebtoonMany(webtoons)
	}
	return err
}
