package controllers

import (
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"hiphopkys/servers/horse/beans"
	"hiphopkys/servers/horse/models"
	"net/http"
)

type SocketmanController struct {
	beego.Controller
}

func (this *SocketmanController) WebSocketJoin() {
	beego.BeeLogger.Error("执行到WebSocketJoin()")
	// Upgrade from http request to WebSocket.
	ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		beego.BeeLogger.Error("SocketmanController::WebSocketJoin(),Not a websocket handshake:%s", err.Error())
		http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		beego.BeeLogger.Error("SocketmanController::WebSocketJoin(),Cannot setup WebSocket connection:%s", err.Error())
		return
	}

	go func(ws *websocket.Conn) {
		defer ws.Close()
		for {
			_, buffer, err := ws.ReadMessage()
			if err != nil {
				beego.BeeLogger.Error("SocketmanController::WebSocketJoin(),读取消息失败:%s", err.Error())
				continue
			}
			bean := beans.ClientRequestBean{}
			err = bean.Unmarshal(buffer)
			if err != nil {
				beego.BeeLogger.Error("SocketmanController::WebSocketJoin(),protobuf解析失败:%s", err.Error())
				continue
			}
			requestId := bean.RequestId
			switch bean.OptionCode {
			case beans.RequestOperationCode_REQUEST_OPERATIONCODE_CREATEROOM: //创建房间
				createRoomBean := bean.GetCreateroomBean()
				beego.BeeLogger.Error("创建房间,requestID:%s,收到的数据:%#v", requestId, createRoomBean)
				go CreateRoomHandler(requestId, ws, createRoomBean)
				// CreateRoomHandler(ws, createRoomBean, requestId)
				break
			case beans.RequestOperationCode_REQUEST_OPERATIONCODE_JOINROOM: //加入房间
				joinRoomBean := bean.GetJoinroomBean()
				beego.BeeLogger.Error("加入房间,requestID:%s,收到的数据:%#v", requestId, joinRoomBean)
				go PlayerJoin(requestId, ws, joinRoomBean)
				break
			case beans.RequestOperationCode_REQUEST_OPERATIONCODE_PLAYERDEVICCEBEAN: //玩家陀螺仪信息
				playerDeviceBean := bean.GetPlayerdeviceBean()
				beego.BeeLogger.Error("玩家陀螺仪信息,requestID:%s,收到的数据:%#v", requestId, playerDeviceBean)
				go BroadcastWebSocket(playerDeviceBean)
				break
			}
		}
	}(ws)

}

func BroadcastWebSocket(bean *beans.PlayerDeviceBean) {
	beego.BeeLogger.Error("收到数据:%#v", bean)
	v := RoomId2PlayingRoomMap.Get(bean.RoomId)
	if nil != v {
		roomModel := (v).(*models.RoomModel)
		beego.BeeLogger.Error("房间信息:%#v", roomModel)
		if buffer, err := bean.Marshal(); err == nil {
			roomModel.Conn.WriteMessage(websocket.BinaryMessage, buffer)
		}
	}
}
