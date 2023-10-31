package main

import (
	"fmt"
	"owl-webtoon/services/webtoon"
)

func main() {
	// webtoon.InsertWebtoonOne(webtoon.Webtoon{Title: "test", Vendor: "kakao"})
	fmt.Println(webtoon.GetWebtoonByTitle("test"))
}
