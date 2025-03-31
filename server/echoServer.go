package server

import (
	"github.com/call-forwarding/config"
	"github.com/call-forwarding/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type echoServer struct {
	app  *echo.Echo
	conf *config.Config
	db   database.Database
}

func NewEchoServer(conf *config.Config, db database.Database) Server {
	echoApp := echo.New()
	echoApp.Logger.SetLevel(log.DEBUG)
	echoApp.Use(middleware.Recover())
	echoApp.Use(middleware.Logger())
	return &echoServer{
		app:  echoApp,
		conf: conf,
		db:   db,
	}
}
