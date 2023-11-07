package router

import (
	"owl-webtoon/validatorAdaptor"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "owl-webtoon/docs"
)

const baseURL = "/api/v1/"

func RouteAll() {
	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Validator = validatorAdaptor.NewParamValidator()
	webtoonRouter(e)
	e.Logger.Fatal(e.Start(":3000"))
}
