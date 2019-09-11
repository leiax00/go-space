package main

import (
	"fmt"
	"go.uber.org/fx"
	"leiax00.com/fxWeb/conf"
	"leiax00.com/fxWeb/fxEcho"
	"leiax00.com/fxWeb/handler"
	"leiax00.com/fxWeb/route"
)

func init() {
	fmt.Println("Hello, FX-WEB!!!!!")
}

func main() {
	ymlFile := "application.yml"
	app := fx.New(
		conf.NewConfModule(&ymlFile),
		fxEcho.NewEchoModule(),
		fx.Provide(
			handler.NewWebHandler,
		),
		fx.Invoke(
			route.RegisterRoutes,
		),
	)
	//ctx, cancel := context.WithTimeout(context.Background(), 30*time.Millisecond)
	//defer cancel()
	//app.Start(ctx)
	app.Run()
}
