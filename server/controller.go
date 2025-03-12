package server

import (
	"fmt"
	"net/http"

	"github.com/call-forwarding/harperDB/handlers"
	"github.com/labstack/echo/v4"
)

func (e *echoServer) Start() {
	// Health check adding
	e.app.GET("v1/health", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	e.app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.app.GET("/hello/:username", func(c echo.Context) error {
		uname := c.Param("username")
		return c.String(http.StatusOK, "Hello, "+uname+"!")
	})
	// handlers.NewHarperHttpHandler(e.db).CfStatus()
	e.app.GET("/call-forward", handlers.NewHarperHttpHandler(e.db).CfStatus)

	serverUrl := fmt.Sprintf(":%d", e.conf.Server.Port)
	e.app.Logger.Fatal(e.app.Start(serverUrl))
}
