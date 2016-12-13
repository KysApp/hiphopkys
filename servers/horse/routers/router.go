package routers

import (
	"github.com/astaxie/beego"
	"hiphopkys/servers/horse/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/game/join", &controllers.SocketmanController{}, "get,post:WebSocketJoin")
	beego.Router("/game/pushAppointmentUser", &controllers.RpcController{}, "get,post:PushAppointmentUser")
}
