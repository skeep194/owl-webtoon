package webtoon

import (
	"owl-webtoon/database"

	"github.com/google/uuid"
)

func GetWebtoonByTitle(title string) Webtoon {
	w := Webtoon{Title: title}

	database.PostgreDB.Where("title = ?", title).First(&w)
	return w
}

func GetWebtoonIdByTitleAndVendor(title string, vendor WebtoonPlatform) uuid.UUID {
	w := Webtoon{}

	database.PostgreDB.Where("title = ? AND vendor = ?", title, vendor).Find(&w)
	return w.ID
}
