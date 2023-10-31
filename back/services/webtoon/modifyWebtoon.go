package webtoon

import (
	"owl-webtoon/database"
)

func InsertWebtoonOne(w Webtoon) {
	database.PostgreDB.Create(&w)
}
