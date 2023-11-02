package router

import (
	"owl-webtoon/controller"

	"github.com/labstack/echo/v4"
)

func webtoonRouter(e *echo.Echo) {
	e.GET(baseURL+"webtoon", controller.GetWebtoon)
}
