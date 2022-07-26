package main

import (
	"fmt"
	"net/http"
	
	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Print("\x1bc")
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1111"))
}
