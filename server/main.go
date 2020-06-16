package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var e = createMux()

func main() {
	e.GET("/", articleIndex)

	e.Logger.Fatal(e.Start(":8080"))
}

// 返り値はechoのポインタ(実態)
func createMux() *echo.Echo {
	// echoを生成
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	return e
}

func articleIndex(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
