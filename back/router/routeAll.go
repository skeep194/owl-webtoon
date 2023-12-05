package router

import (
	"owl-webtoon/validatorAdaptor"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "owl-webtoon/docs"
)

const baseURL = "/api/v1/"

func RouteAll() {
	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Validator = validatorAdaptor.NewParamValidator()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
	}))
	webtoonRouter(e)
	e.Logger.Fatal(e.Start(":3000"))
}
