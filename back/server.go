package main

import (
	"owl-webtoon/cron"
	"owl-webtoon/router"
)

func main() {
	cron.CrawlPerOneDay()
	router.RouteAll()
}
