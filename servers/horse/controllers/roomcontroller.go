package controllers

import (
	"container/list"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"hiphopkys/servers/horse/beans"
	"hiphopkys/servers/horse/models"
)

var (
	WaitPlayerJoinRoomList  = list.New() //未满,正在等待玩家加入的房间队列
	FullWaitPlayingRoolList = list.New() //已满,等待游戏开始的房间队列
	PlayingRoomList         = list.New() //已满,游戏正在进行中的房间队列
)

var (
	JoinWaitPlayRoomChan        = make(chan *models.RoomModel, 500) //未满,正在等待玩家加入的房间队列管道
	JoinFullWaitPlayingRoomChan = make(chan *models.RoomModel, 500) //已满,等待游戏开始的房间队列管道
	JoinPlayingRoomChan         = make(chan *models.RoomModel, 500) //已满,游戏正在进行中的房间队列管道
)

func init() {
	go loop()
}

/**
 * 加入WaitPlayerJoinRoomList
 * @param {[type]} requestId string                [请求Id,PC端请求新疆房间时有客户端生成并传递的]
 * @param {[type]} conn      *websocket.Conn       [与PC端的websocket链接]
 * @param {[type]} bean      *beans.CreateRoomBean [PC端请求创建房间时传递的Bean]
 */
func CreateRoomHandler(requestId string, conn *websocket.Conn, bean *beans.CreateRoomBean) {
	roomModel := models.NewRoomModel(conn, bean.GameId, bean.Longitude, bean.Latitude, bean.DeviceInfo)
	WaitPlayerJoinRoomList <- roomModel
}

func loop() {
	for {
		select {
		case roomModel := <-JoinWaitPlayRoomChan:
			WaitPlayerJoinRoomList.PushBack(roomModel)
			break
		case roomModel := <-JoinFullWaitPlayingRoomChan:
			break
		case roomModel := <-JoinPlayingRoomChan:
			break
		}
	}
}
