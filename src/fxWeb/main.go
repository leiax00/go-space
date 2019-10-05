package main

import (
	"context"
	"github.com/labstack/gommon/log"
	"go.uber.org/fx"
	"leiax00.com/fxWeb/conf"
	"leiax00.com/fxWeb/fxEcho"
	"leiax00.com/fxWeb/web"
)

func init() {
	log.Info("Hello, FX-WEB!!!!!")
}

func main() {
	ymlFile := "application.yml"
	newApp(ymlFile).Run()

}

func newApp(ymlFile string) *fx.App {
	return fx.New(
		conf.NewConfModule(&ymlFile),
		fxEcho.NewEchoModule(),
		web.NewModule(),
		fx.Provide(
		//handler.NewWebHandler,
		),
		fx.Invoke(
			registerHook,
			//web.RegisterRoutes,
		),
		fx.ErrorHook(&fwErrorHandler{}),
	)
}

func registerHook(lc fx.Lifecycle) {
	lc.Append(fx.Hook{
		OnStart: startHook,
		OnStop:  stopHook,
	})
}

func startHook(ctx context.Context) error {
	fxEcho.StartFxEcho()
	return nil
}

func stopHook(ctx context.Context) error {
	log.Info("FX-WEB:Goodbye~~")
	return nil
}

type fwErrorHandler struct{}

func (f fwErrorHandler) HandleError(err error) {
	log.Error("Failed to start container, err:", err.Error())
}
