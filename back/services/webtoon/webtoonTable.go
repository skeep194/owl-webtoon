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
	Likes int       `gorm:"default:0"`

	Vendor     WebtoonPlatform `gorm:"type:vendor;index:webtoon,priority:2"`
	SiteID     string
	Title      string `gorm:"index:webtoon,priority:1"`
	IsVacation bool
	//data represent as bitmask monday = 2^0, .. sunday = 2^6, irregular = 2^7 if all bit has zero(we assume that value is 0), webtoon is complete and no longer update.
	UploadDate   uint8
	CreateDate   time.Time
	UpdateDate   time.Time
	IsComplete   bool
	EpisodeCount int
	LinkDomain   string
}

func init() {
	database.PostgreDB.AutoMigrate(&Webtoon{})
}
