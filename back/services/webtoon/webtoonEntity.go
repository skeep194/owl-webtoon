package webtoon

import (
	"owl-webtoon/database"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
)

type WebtoonPlatform int

const (
	Naver = WebtoonPlatform(iota)
	Kakao
)

type DayName int

const (
	Mon = DayName(iota)
	Tue
	Wed
	Thu
	Fri
	Sat
	Sun
)

type WebtoonEntity struct {
	Platform  WebtoonPlatform
	Title     string
	ContentId string
	Day       []DayName
	Link      string
}

func getWebtoonCollection() (db *mongo.Collection, err error) {
	db, err = database.GetCollection("webtoon")
	if err != nil {
		return nil, errors.Wrap(err, "fail to get webtoon collection")
	}
	return db, nil
}
