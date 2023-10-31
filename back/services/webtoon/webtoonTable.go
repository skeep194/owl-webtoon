package webtoon

import (
	"owl-webtoon/database"
	"time"

	"github.com/google/uuid"
)

type WebtoonPlatform string

const (
	Naver WebtoonPlatform = "naver"
	Kakao WebtoonPlatform = "kakao"
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

type Webtoon struct {
	ID    uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Likes int

	Vendor       WebtoonPlatform `gorm:"type:vendor"`
	SiteID       string
	Title        string
	IsVacation   bool
	UploadDate   string
	CreateDate   time.Time
	UpdateDate   time.Time
	IsComplete   bool
	EpisodeCount int
	LinkDomain   string
}

func init() {
	database.PostgreDB.AutoMigrate(&Webtoon{})
}
