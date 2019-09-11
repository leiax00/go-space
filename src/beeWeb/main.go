package main

import (
	"github.com/astaxie/beego"
	_ "github.com/beeWeb/routers"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		http.ListenAndServe("0.0.0.0:8090", nil)
	}()
	beego.Run()
}
