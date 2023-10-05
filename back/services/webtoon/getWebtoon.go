package webtoon

import (
	"context"
)

func InsertWebtoonEntity(w WebtoonEntity) error {
	db, err := getWebtoonCollection()
	if err != nil {
		return err
	}
	_, err = db.InsertOne(context.TODO(), w)
	return err
}
