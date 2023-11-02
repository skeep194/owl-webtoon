package webtoonService

import (
	"owl-webtoon/database"

	"github.com/google/uuid"
)

func GetWebtoons(title string, vendor WebtoonPlatform) (ret []Webtoon) {
	w := Webtoon{Title: title, Vendor: vendor}

	database.PostgreDB.Where(&w).Find(&ret)
	return ret
}

func getWebtoonIdByTitleAndVendor(title string, vendor WebtoonPlatform) uuid.UUID {
	w := Webtoon{}

	database.PostgreDB.Where("title = ? AND vendor = ?", title, vendor).Find(&w)
	return w.ID
}
