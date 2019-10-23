package fxEcho

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go.uber.org/fx"
	"net/http"
	"os"
	"time"
)

var e *echo.Echo

func NewEchoModule() fx.Option {
	return fx.Options(fx.Provide(newFxEcho))
}

func newFxEcho() *echo.Echo {
	e = echo.New()
	setConfig()
	return e
}

func StartFxEcho() {
	s := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}
	go func() {
		e.Logger.Fatal(e.StartServer(s))
		//go e.Start(":8080")
	}()
}

func setConfig() {
	if l, ok := e.Logger.(*log.Logger); ok {
		l.SetHeader("${time_rfc3339} ${level}")
		l.SetOutput(os.Stdout)
		l.SetLevel(log.DEBUG)
	}

	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &FxEchoContext{c}
			return h(cc)
		}
	})
}
