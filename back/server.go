package main

import (
	"fmt"
	"owl-webtoon/services/webtoon"
)

func main() {
	// fmt.Println(dataCrawl.GetWeekdayWebtoonURLs())
	err := webtoon.InsertWebtoonEntity(webtoon.WebtoonEntity{
		Platform:  webtoon.Naver,
		Title:     "레후",
		ContentId: "콘텐트아이디레후",
		Day:       []webtoon.DayName{webtoon.Fri},
		Link:      "www.refu.regu",
	})
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
}
