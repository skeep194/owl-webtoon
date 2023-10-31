package webtoon

import "owl-webtoon/database"

func GetWebtoonByTitle(title string) Webtoon {
	w := Webtoon{Title: title}

	database.PostgreDB.Where("title = ?", title).First(&w)
	return w
}
