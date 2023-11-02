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

func GetWebtoon(c echo.Context) (err error) {
	getWebtoonParam := new(GetWebtoonParam)
	if err = c.Bind(getWebtoonParam); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	ret := webtoonService.GetWebtoons(getWebtoonParam.Title, getWebtoonParam.Vendor)
	return c.JSON(http.StatusOK, ret)
}
