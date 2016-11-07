package controllers

import (
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"hiphopkys/servers/commons/container/list"
	"hiphopkys/servers/commons/container/smap"
	"hiphopkys/servers/horse/beans"
	"hiphopkys/servers/horse/models"
	"sync"
)

var (
	StorageRoomInstanceChan           = make(chan *models.RoomModel, 500) //存储RoomModel实体Channel
	InitRoomChan                      = make(chan string, 500)            //初始化房间channel,等待分配类型
	PrepareStartGameRoomChan          = make(chan string, 500)            //房间已满，等待游戏前的初始化工作的房间channel
	AppointmentWaitCheckStateRoomChan = make(chan string, 500)            //预约玩家房间临时channel
	AppointmentWaitJoinRoomChan       = make(chan string, 500)            //预约玩家排队等待进入房间Channel(高优先级)
)

var (
	RoomId2RoomInstanceMap = smap.New() //记录所有的房间实体
	WaitPlayerJoinRoomList = list.New() //等待玩家(预约、非预约)加入房间队列
	PlayingRoomList        = list.New() //正在游戏的房间队列
)

func init() {
	go loop_storage_roomInstance()
	go loop_room_dispatch()
	go loop_roomstate_update()
	go loop_prepare_startgame()
}

/**
 * 加入WaitPlayerJoinRoomList
 * @param {[type]} requestId string                [请求Id,PC端请求新疆房间时有客户端生成并传递的]
 * @param {[type]} conn      *websocket.Conn       [与PC端的websocket链接]
 * @param {[type]} bean      *beans.CreateRoomBean [PC端请求创建房间时传递的Bean]
 */
func CreateRoomHandler(requestId string, conn *websocket.Conn, bean *beans.CreateRoomBean) {
	roomModel := models.NewRoomModel(conn, bean.GameId, bean.Longitude, bean.Latitude, bean.DeviceInfo)
	StorageRoomInstanceChan <- roomModel
}

/**
 * 存储用RoomModel实体
 * @return {[type]} [description]
 */
func loop_storage_roomInstance() {
	for {
		select {
		case roomModel := <-StorageRoomInstanceChan:
			if nil != roomModel {
				RoomId2RoomInstanceMap.Insert(roomModel.RoomId, roomModel)
				InitRoomChan <- roomModel.RoomId
			}
		}
	}
}

/**
 *	预约玩家房间调度
 * @return {[type]} [description]
 */
func loop_room_dispatch() {
	for {
		select {
		case playerID := <-AppointmentWaitJoinPlayerChan:
			select {
			case roomID := <-InitRoomChan:
				/**
				 *此时说明有预约玩家等待加入房间并且有房间建立,
				 *则需要 1.将该玩家加入房间，2.初始化房间，包括设置房间类型为预约玩家房间和设置预约ID 3.将房间加入WaitPlayerJoinRoomList
				 */
				roomModel := (RoomId2RoomInstanceMap.Get(roomID)).(*models.RoomModel)
				roomModel.PlayerUserIdList.PushBack(roomModel)
				playerModel := (PlayerId2RoomInstanceMap.Get(playerID)).(*models.PlayerModel)
				roomModel.AppointmentId = playerModel.AppointmentId
				WaitPlayerJoinRoomList.PushBack(roomID)
			}
		}
	}
}

/**
 * 房间状态更新模块,从预约玩家房间临时channel取出一个进行检查是否可以开始游戏
 * @return {[type]} [description]
 */
func loop_roomstate_update() {
	for {
		select {
		case roomID := <-AppointmentWaitCheckStateRoomChan:
			roomModel := (RoomId2RoomInstanceMap.Get(roomID)).(*models.RoomModel)
			if roomModel.Capacity == int32(roomModel.PlayerUserIdList.Len()) { //房间已满可以开始游戏,加入初始化队列
				PrepareStartGameRoomChan <- roomID
			}
			/**
			 * 通知其他玩家现在房间中的玩家playerid list(可选,可以等到最后所有玩家到齐后准备开始游戏时统一返回)
			 */
		}
	}
}

func loop_prepare_startgame() {
	for {
		select {
		case roomID := <-PrepareStartGameRoomChan:
			WaitPlayerJoinRoomList.RemoveFirstElementWithValue(roomID)

		}
	}
}
