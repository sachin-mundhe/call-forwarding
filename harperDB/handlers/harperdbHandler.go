package handlers

import "github.com/labstack/echo/v4"

type HarperdbHandler interface {
	UnConditionalCF(c echo.Context) error
	CallForwardings(c echo.Context) error
}
