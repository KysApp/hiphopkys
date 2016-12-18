package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"hiphopkys/servers/commons/container/smap"
	"hiphopkys/servers/horse/beans"
	"hiphopkys/servers/horse/models"
	"net/http"
)

var (
	Conn2PlayerIdSmap = smap.New()
	Conn2RoomIdSmap   = smap.New()
)

type SocketmanController struct {
	beego.Controller
}

func (this *SocketmanController) WebSocketJoin() {
	beego.BeeLogger.Error("执行到WebSocketJoin()")
	// Upgrade from http request to WebSocket.
	responseBean := &beans.RPCResponse{}
	responseBean.ErrorCode = "0"
	responseBean.Desc = "success"
	ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 2048, 2048)
	if _, ok := err.(websocket.HandshakeError); ok {
		beego.BeeLogger.Error("SocketmanController::WebSocketJoin(),Not a websocket handshake:%s", err.Error())
		http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		responseBean.ErrorCode = beego.AppConfig.String("errcode::websocket_nothandler_error_code")
		responseBean.Desc = beego.AppConfig.String("errcode::websocket_nothandler_error_desc")
	} else if err != nil {
		beego.BeeLogger.Error("SocketmanController::WebSocketJoin(),Cannot setup WebSocket connection:%s", err.Error())
		responseBean.ErrorCode = beego.AppConfig.String("errcode::websocket_conn_error_code")
		responseBean.Desc = beego.AppConfig.String("errcode::websocket_conn_error_desc")
	}

	go func(ws *websocket.Conn) {
		defer ws.Close()
		for {
			_, buffer, err := ws.ReadMessage()
			if err != nil {
				beego.BeeLogger.Error("SocketmanController::WebSocketJoin(),读取消息失败:%s", err.Error())
				if websocket.IsCloseError(err, websocket.CloseNormalClosure) ||
					websocket.IsCloseError(err, websocket.CloseGoingAway) ||
					websocket.IsCloseError(err, websocket.CloseProtocolError) ||
					websocket.IsCloseError(err, websocket.CloseUnsupportedData) ||
					websocket.IsCloseError(err, websocket.CloseNoStatusReceived) ||
					websocket.IsCloseError(err, websocket.CloseAbnormalClosure) ||
					websocket.IsCloseError(err, websocket.CloseInvalidFramePayloadData) ||
					websocket.IsCloseError(err, websocket.CloseMessageTooBig) ||
					websocket.IsCloseError(err, websocket.CloseServiceRestart) ||
					websocket.IsCloseError(err, websocket.CloseTryAgainLater) ||
					websocket.IsCloseError(err, websocket.CloseTLSHandshake) {
					break
				}
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

	buf, _ := json.Marshal(responseBean)
	this.Ctx.WriteString(string(buf))
}

func BroadcastWebSocket(bean *beans.PlayerDeviceBean) {
	beego.BeeLogger.Error("收到数据:%#v", bean)
	wsResponseBean := beans.ServerSendBean{}
	wsResponseBean.OptionCode = beans.SendMessageOperationCode_SENDMESSAGE_OPERATIONCODE_PLAYERDEVICE
	wsResponseBean.Bean = &beans.ServerSendBean_PlayerDeviceBean{
		PlayerDeviceBean: bean,
	}
	if msgBuf, err := wsResponseBean.Marshal(); nil == err {
		v := RoomId2PlayingRoomMap.Get(bean.RoomId)
		if nil != v {
			roomModel := (v).(*models.RoomModel)
			beego.BeeLogger.Error("房间信息:%#v", roomModel)
			roomModel.Conn.WriteMessage(websocket.BinaryMessage, msgBuf)
		}
	} else {
		beego.BeeLogger.Error("BroadcastWebSocket中protobuf转换错误:%s", err.Error())
	}

}
