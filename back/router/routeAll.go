package router

import (
	"owl-webtoon/validatorAdaptor"

	"github.com/labstack/echo/v4"
)

const baseURL = "/api/v1/"

func RouteAll() {
	e := echo.New()
	e.Validator = validatorAdaptor.NewParamValidator()
	webtoonRouter(e)
	e.Logger.Fatal(e.Start(":3000"))
}
