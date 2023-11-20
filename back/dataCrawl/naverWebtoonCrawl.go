package dataCrawl

import (
	"fmt"
	"net/url"
	"owl-webtoon/services/webtoonService"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
)

type naverWebtoonCrawller struct {
	*staticCrawller
}

type naverWebtoonType int

const (
	weekday naverWebtoonType = iota
	complete
	dailyPlus
)

const baseURL string = "https://m.comic.naver.com"

// webtoon weekday or finish page to webtoon URL.
// example path
// weekday?week=wed
func (naver naverWebtoonCrawller) getWebtoonURLs(path string) (ret []string, err error) {
	naver.staticCrawller, err = newStaticCrawller(baseURL + "/webtoon/" + path)
	if err != nil {
		return nil, errors.Wrap(err, "fail to make static crawller "+path)
	}
	defer naver.staticCrawller.closeFunc()
	naver.document.Selection.Find("div.section_list_toon a.link").Each(func(idx int, sel *goquery.Selection) {
		str, isExist := sel.Attr("href")
		if isExist {
			ret = append(ret, str)
		}
	})
	return ret, nil
}

// webtoon main page to webtoon instance.
// example path
// list?titleId=26460
func (naver naverWebtoonCrawller) getWebtoonInfo(path string, webtoonType naverWebtoonType) (ret webtoonService.Webtoon, err error) {
	URL, err := url.Parse(baseURL + path)
	if err != nil {
		return ret, errors.Wrap(err, "invalid url")
	}

	dayString := map[string]uint8{"월": 1 << 0, "화": 1 << 1, "수": 1 << 2, "목": 1 << 3, "금": 1 << 4, "토": 1 << 5, "일": 1 << 6, "매일": (1 << 7) - 1}

	naver.staticCrawller, err = newStaticCrawller(baseURL + path + "&sortOrder=DESC")
	if err != nil {
		return ret, errors.Wrap(err, "fail to make static crawller "+path+"&sortOrder=DESC")
	}
	defer naver.staticCrawller.closeFunc()
	var find string

	//set SiteID
	m, _ := url.ParseQuery(URL.RawQuery)
	ret.SiteID = m["titleId"][0]

	//set Vendor
	ret.Vendor = webtoonService.Naver

	//set Title
	ret.Title = naver.document.Selection.Find("div.area_info strong.title").Text()

	//set IsVacation
	ret.IsVacation = naver.document.Selection.Find("div.area_info span.bullet.break").Length() != 0

	//set UploadDate
	switch webtoonType {
	case complete:
		ret.UploadDate = 0
	case weekday:
		naver.document.Selection.Find("ul.list_detail li").Each(func(i int, s *goquery.Selection) {
			ret.UploadDate |= dayString[s.Text()]
		})
	case dailyPlus:
		ret.UploadDate = 1 << 7
	default:
		err = errors.New("unexpected error webtoonType invalid")
	}
	if err != nil {
		return ret, err
	}

	//set UpdateDate
	find = naver.document.Selection.Find("li.item span.date").First().Text()
	ret.UpdateDate, err = time.Parse("06.01.02", find)
	if err != nil {
		return ret, errors.Wrap(err, "date parse err "+find)
	}

	//set IsComplete
	ret.IsComplete = webtoonType == complete

	//set EpisodeCount
	find = naver.document.Selection.Find("h3.title_count span.count_num").Text()
	ret.EpisodeCount, err = strconv.Atoi(find)
	if err != nil {
		return ret, errors.Wrap(err, "convert to int "+find)
	}

	//set LinkDomain
	ret.LinkDomain = path

	//set CreateDate
	naver.staticCrawller, err = newStaticCrawller(baseURL + path + "&sortOrder=ASC")
	if err != nil {
		return ret, errors.Wrap(err, "fail to make static crawller "+path+"&sortOrder=ASC")
	}
	defer naver.staticCrawller.closeFunc()
	find = naver.document.Selection.Find("li.item span.date").First().Text()
	ret.CreateDate, err = time.Parse("06.01.02", find)
	if err != nil {
		return ret, errors.Wrap(err, "date parse err "+find)
	}
	return ret, nil
}

func (naver naverWebtoonCrawller) getFinishPageCount() (ret int, err error) {
	naver.staticCrawller, err = newStaticCrawller(baseURL + "/webtoon/finish")
	if err != nil {
		return -1, errors.Wrap(err, "fail to make static crawller finish")
	}
	defer naver.staticCrawller.closeFunc()
	find := naver.document.Selection.Find("em.current_pg span.total").Text()
	ret, err = strconv.Atoi(find)
	if err != nil {
		return -1, errors.Wrap(err, "convert to int "+find)
	}
	return ret, nil
}

// TODO: need error logging.
// temporarily error is ignored
func (naver naverWebtoonCrawller) getWebtoonInfos() (ret []webtoonService.Webtoon, err error) {
	finishPageCount, err := naver.getFinishPageCount()
	fmt.Printf("finish page %d", finishPageCount)
	var paths = []string{"mon", "tue", "wed", "thu", "fri", "sat", "sun", "dailyPlus"}
	dayCount := len(paths)
	for i := 1; i <= finishPageCount; i++ {
		paths = append(paths, fmt.Sprintf("finish?page=%d&sort=UPDATE", i))
	}
	//get naver webtoon URL from mobile site
	var URLs = []string{}
	var webtoonTypes = []naverWebtoonType{}
	for i, path := range paths {
		if i < dayCount {
			path = "weekday?week=" + path
		}

		var tmp []string
		tmp, err = naver.getWebtoonURLs(path)
		if err != nil {
			err = errors.Wrap(err, "fail to get webtoon urls "+path)
			continue
		}
		var webtoonType naverWebtoonType
		switch {
		case paths[i] == "dailyPlus":
			webtoonType = dailyPlus
		case i < dayCount:
			webtoonType = weekday
		default:
			webtoonType = complete
		}
		URLs = append(URLs, tmp...)
		for j := 0; j < len(tmp); j++ {
			webtoonTypes = append(webtoonTypes, webtoonType)
		}
	}
	//get webtoon information from all URL
	for i, URL := range URLs {
		var w webtoonService.Webtoon
		w, err = naver.getWebtoonInfo(URL, webtoonTypes[i])
		if err != nil {
			err = errors.Wrap(err, "fail to get webtoon info "+URL)
			continue
		}
		fmt.Printf("%dth success\n", i)
		ret = append(ret, w)
	}
	return ret, err
}
