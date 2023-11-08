package main

import (
	"fmt"
	"owl-webtoon/cron"
	"owl-webtoon/router"
)

// @title owl-webtoon API
// @version 1.0
// @description 부엉이는 부엉부엉하고 울어요

// @host 34.64.107.125
// @BasePath /api/v1
func main() {
	fmt.Println("server start!!!")
	cron.CrawlPerOneDay()
	router.RouteAll()
}
