package dataCrawl

import (
	"fmt"
	"owl-webtoon/services/webtoonService"

	"github.com/go-co-op/gocron"
	"github.com/pkg/errors"
)

var crawllers = []Crawller{
	naverWebtoonCrawller{},
}

func Crawl(job gocron.Job) (err error) {
	fmt.Printf("crawl last run: %s crawl next run: %s\n", job.LastRun(), job.NextRun())
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
