package dataCrawl

import (
	"owl-webtoon/services/webtoon"

	"github.com/pkg/errors"
)

var crawllers = []Crawller{
	naverWebtoonCrawller{},
}

func Crawl() (err error) {
	for _, crawller := range crawllers {
		var webtoons []webtoon.Webtoon
		webtoons, err = crawller.getWebtoonInfos()
		if err != nil {
			err = errors.Wrap(err, "some error from get webtoon infos")
		}
		webtoon.InsertWebtoonMany(webtoons)
	}
	return err
}
