package webtoon

import (
	"owl-webtoon/database"
)

func InsertWebtoonMany(webtoons []Webtoon) {
	for _, webtoon := range webtoons {
		webtoon.ID = getWebtoonIdByTitleAndVendor(webtoon.Title, webtoon.Vendor)
		database.PostgreDB.Save(&webtoon)
	}
}
