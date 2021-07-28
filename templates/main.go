package main

import (
	"net/http"
	_ "net/http/pprof"

	router "gitlab.com/zendrulat123/echojw/routing"

	temp "gitlab.com/zendrulat123/echojw/rendor"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Renderer = temp.Rend()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))
	router.Routing(e)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/static", "static")
	e.Logger.Fatal(e.Start(":1323"))
}
