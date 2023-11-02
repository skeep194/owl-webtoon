package main

import (
	"owl-webtoon/dataCrawl"
)

func main() {
	// webtoon.InsertWebtoonOne(webtoon.Webtoon{Title: "test", Vendor: "kakao"})
	// fmt.Println(webtoon.GetWebtoonIdByTitleAndVendor("ad", webtoon.Kakao))
	dataCrawl.Crawl()
}
