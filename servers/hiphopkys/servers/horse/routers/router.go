package routers

import (
	"hiphopkys/servers/hiphopkys/servers/horse/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
