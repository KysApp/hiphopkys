package main

import (
	"flag"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"hiphopkys/servers/mq/consumer/horse/beans"
	_ "hiphopkys/servers/mq/consumer/horse/handlers"
	"log"
	"net/http"
)

var addr = flag.String("addr", "localhost:7475", "http service address")
var upgrader = websocket.Upgrader{}

func handler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		beego.BeeLogger.Error("upgrade:%s", err.Error())
		return
	}
	defer ws.Close()
	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			beego.BeeLogger.Error("read:%s", err.Error())
			break
		}
		reciveBean := beans.User{}
		reciveBean.Unmarshal(message)
		// beego.BeeLogger.Error("recv: %#v", reciveBean)
		log.Printf("recv: %#v", reciveBean)
		err = ws.WriteMessage(websocket.TextMessage, []byte("收到"))
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
