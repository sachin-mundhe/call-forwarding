package handlers

import "github.com/labstack/echo/v4"

type HarperdbHandler interface {
	CfStatus(c echo.Context) error
}
