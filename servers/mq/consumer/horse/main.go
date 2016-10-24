package main

import (
	"flag"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	_ "hiphopkys/servers/mq/consumer/horse/handlers"
	"net/http"
)

var addr = flag.String("addr", "localhost:8888", "http service address")
var upgrader = websocket.Upgrader{}

func handler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		beego.BeeLogger.Error("upgrade:%s", err.Error())
		return
	}
	defer ws.Close()
	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			beego.BeeLogger.Error("read:%s", err.Error())
			break
		}
		beego.BeeLogger.Error("recv: %s", message)
		err = ws.WriteMessage(mt, message)
		if err != nil {
			beego.BeeLogger.Error("write:%s", err.Error())
			break
		}
	}
}

func main() {
	flag.Parse()
	http.HandleFunc("/horse", handler)
	err := http.ListenAndServe(*addr, nil)
	if nil != err {
		beego.BeeLogger.Error("消息队列消费者(horse)监听失败:%s", err.Error())
	}
}
