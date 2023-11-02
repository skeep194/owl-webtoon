package webtoon

import (
	"owl-webtoon/database"
)

func InsertWebtoonMany(webtoons []Webtoon) {
	for _, webtoon := range webtoons {
		database.PostgreDB.Save(&webtoon)
	}
}
