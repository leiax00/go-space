package routers

import (
	"github.com/astaxie/beego"
	"github.com/beeWeb/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
