package fxEcho

import (
	"github.com/labstack/echo/v4"
)

type FxEchoContext struct {
	echo.Context
}

func (fc *FxEchoContext) Pre() *FxEchoContext {
	fc.Logger().Info("6666666666")
	return fc
}
