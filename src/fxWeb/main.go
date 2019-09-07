package main

import (
	"fmt"
	"go.uber.org/fx"
	"leiax00.com/fxWeb/basic"
	"leiax00.com/fxWeb/conf"
	"leiax00.com/fxWeb/handler"
)

func init() {
	fmt.Println("Hello, FX-WEB!!!!!")
}

func main() {
	app := fx.New(
		basic.NewEchoModule(),
		fx.Provide(
			conf.NewConfig,
			handler.NewWebHandler,
		),
		fx.Invoke(
			basic.RegisterRoutes,
		),
	)
	//ctx, cancel := context.WithTimeout(context.Background(), 30*time.Millisecond)
	//defer cancel()
	//app.Start(ctx)
	go basic.EchoObj.Start(":8080")
	app.Run()
}
