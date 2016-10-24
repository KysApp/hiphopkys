package routers

import (
	"hiphopkys/servers/horse/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
