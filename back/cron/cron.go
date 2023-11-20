package cron

import (
	"fmt"
	"owl-webtoon/dataCrawl"
	"time"

	"github.com/go-co-op/gocron"
)

func CrawlPerOneDay() {
	fmt.Println("cron")
	scheduler := gocron.NewScheduler(time.Local)
	scheduler.SingletonModeAll()
	scheduler.Every(1).Days().DoWithJobDetails(dataCrawl.Crawl)
	scheduler.StartAsync()
}
