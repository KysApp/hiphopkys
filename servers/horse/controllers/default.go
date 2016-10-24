package controllers

import (
	"github.com/astaxie/beego"
	"hiphopkys/servers/mq/consumer"
	"hiphopkys/servers/mq/producer"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	consumer.PrintSelf()
	producer.PrintSelf()
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
