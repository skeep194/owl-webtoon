package cron

import (
	"owl-webtoon/dataCrawl"
	"time"

	"github.com/go-co-op/gocron"
)

func CrawlPerOneDay() {
	scheduler := gocron.NewScheduler(time.Local)
	scheduler.Every(1).Day().At("00:30").Do(dataCrawl.Crawl)
	scheduler.StartAsync()
}
