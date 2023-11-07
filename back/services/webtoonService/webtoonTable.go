package webtoonService

import (
	"io/ioutil"
	"os"
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

// Webtoon model info
// @Description Webtoon information
type Webtoon struct {
	ID uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"-"`
	//부엉이 웹툰에서의 좋아요 수(각 웹툰사와 관계없음)
	Likes int `gorm:"default:0" json:"likes" example:"2"`

	//웹툰 배급사
	Vendor WebtoonPlatform `gorm:"type:vendor;index:webtoon,priority:2" json:"vendor" example:"naver"`
	//각 웹툰 사이트가 자체적으로 부여한 웹툰 id값
	SiteID string `json:"site_id" example:"654774"`
	//웹툰 제목
	Title string `gorm:"index:webtoon,priority:1" json:"title" example:"소녀의 세계"`
	//휴재 중인 경우 true
	IsVacation bool `json:"is_vacation" example:"false"`
	//data represent as bitmask monday = 2^0, .. sunday = 2^6, irregular = 2^7 if all bit has zero(we assume that value is 0), webtoon is complete and no longer update.
	UploadDate uint8 `json:"upload_date" example:"1"`
	//1화 연재 날짜
	CreateDate time.Time `json:"create_date" example:"2015-05-17T09:00:00+09:00"`
	//최신화 업데이트 날짜
	UpdateDate time.Time `json:"update_date" example:"2023-10-29T09:00:00+09:00"`
	//완결 여부
	IsComplete bool `json:"is_complete" example:"false"`
	//총 화수
	EpisodeCount int `json:"episode_count" example:"391"`
	//baseURL을 제외한 웹툰 링크
	LinkDomain string `json:"link_domain" example:"/webtoon/list?titleId=654774\u0026week=mon"`
}

func init() {
	//TODO: sql 파일이 여러개 있을 때도 처리해야함
	sql, _ := os.Open("./services/webtoonService/vendor_enum.sql")
	data, _ := ioutil.ReadAll(sql)
	s := string(data[:])

	database.PostgreDB.Exec(s)
	database.PostgreDB.AutoMigrate(&Webtoon{})
}
