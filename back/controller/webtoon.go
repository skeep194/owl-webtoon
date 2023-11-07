package controller

import (
	"net/http"
	"owl-webtoon/services/webtoonService"

	"github.com/labstack/echo/v4"
)

type GetWebtoonParam struct {
	Title  string                         `query:"title"`
	Vendor webtoonService.WebtoonPlatform `query:"vendor"`
}

// @Summary Get webtoon information
// @Description 크롤링된 웹툰 정보를 보여줍니다. title 혹은 vendor 가 있으면 매칭되는 웹툰 정보만 보여줍니다.
// @Accept json
// @Produce json
// @Param title query string false "title 일치하는 웹툰 정보"
// @Param vendor query webtoonService.WebtoonPlatform false "배급사가 vendor인 웹툰들 정보" Enum("naver", "kakao")
// @Success 200 {array} webtoonService.Webtoon "list of matched webtoon"
// @Router /webtoon [get]
func GetWebtoon(c echo.Context) (err error) {
	getWebtoonParam := new(GetWebtoonParam)
	if err = c.Bind(getWebtoonParam); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	ret := webtoonService.GetWebtoons(getWebtoonParam.Title, getWebtoonParam.Vendor)
	return c.JSON(http.StatusOK, ret)
}
