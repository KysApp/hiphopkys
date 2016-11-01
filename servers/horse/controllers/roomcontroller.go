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
	PlayingEndRoomList      = list.New() //已满,一局游戏结束
)

var (
	PrepareStartGameChan = make(byte, 500) //游戏开始前的准备信号
	NowStartGameChan     = make(byte, 500) //游戏开始信号
)

// var (
// 	NewPlayerJoinRoomChan = make()
// )

var (
	JoinWaitPlayRoomChan        = make(chan *models.RoomModel, 500) //未满,正在等待玩家加入的房间队列管道
	JoinFullWaitPlayingRoomChan = make(chan *models.RoomModel, 500) //已满,等待游戏开始的房间队列管道
	JoinPlayingRoomChan         = make(chan *models.RoomModel, 500) //已满,游戏正在进行中的房间队列管道
	JoinPlayingEndRoomChan      = make(chan *models.RoomModel, 500) //已满,一局游戏结束管道
)

func init() {
	go loop()
	go loop_prepareStartGame()
	go loop_nowStartGame()
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

/**
 * 从JoinFullWaitPlayingRoomChan取出第一个房间做一些游戏开始前的初始化工作
 * @return {[type]} [description]
 */
func loop_prepareStartGame() {
	for {
		select {
		case <-PrepareStartGameChan:
			frontElement := FullWaitPlayingRoolList.Front()
			if nil == frontElement {
				beego.BeeLogger.Error("FullWaitPlayingRoolList为空")
				break
			}
			roomModel := frontElement.Value.(*models.RoomModel)
			if nil == roomModel {
				beego.BeeLogger.Error("FullWaitPlayingRoolList中元素数据有误")
				break
			}
			/**************************************
			***************************************
			***这里开始游戏开始前对房间的初始化工作*****
			****************************************
			****************************************/
			//准备工作完毕后从FullWaitPlayingRoolList删除
			FullWaitPlayingRoolList.Remove(frontElement)
			JoinPlayingRoomChan <- roomModel
		}
	}
}

/**
 * Server -> client
 * 通知PC端以及Player手机端开始游戏，这时候开始接收摇一摇的数据
 * @return {[type]} [description]
 */
func loop_nowStartGame() {
	for {
		select {
		case <-NowStartGameChan:
			frontElement := PlayingRoomList.Front()
			if nil == frontElement {
				beego.BeeLogger.Error("PlayingRoomList为空")
				break
			}
			roomModel := frontElement.Value.(*models.RoomModel)
			if nil == roomModel {
				beego.BeeLogger.Error("PlayingRoomList中元素数据有误")
				break
			}
			/**************************************
			***************************************
			***这里开始通知客户端房间已满可以开始游戏*****
			****************************************
			****************************************/
		}
	}
}

/**
 *	Loop,负责当房间状态改变(房间新建,房间已满但还未开始游戏,房间已满并且游戏正在进行中,房间已满并且一局游戏结束)时的Process
 * @return {[type]} [description]
 */
func loop() {
	for {
		select {
		case roomModel := <-JoinWaitPlayRoomChan: //未满,正在等待玩家加入的房间队列
			WaitPlayerJoinRoomList.PushBack(roomModel)
			roomModel.RoomState = models.ROOM_STATE_WAITJOIN
			break
		case roomModel := <-JoinFullWaitPlayingRoomChan: //已满,等待游戏开始的房间队列,这里只负责添加，删除部分需要信号传递方处理
			FullWaitPlayingRoolList.PushBack(roomModel)
			roomModel.RoomState = models.ROOM_STATE_FULL_WAIT_PLAYING
			PrepareStartGameChan <- '0' //发送消息给loop_prepareStartGame，完成游戏开始前的准备工作
			break
		case roomModel := <-JoinPlayingRoomChan: //已满,游戏正在进行中的房间队列
			PlayingRoomList.PushBack(roomModel)
			roomModel.RoomState = models.ROOM_STATE_PLAYING
			NowStartGameChan <- '0'
			break
		case roomModel := <-JoinPlayingEndRoomChan: //已满,一局游戏结束
			PlayingEndRoomList.PushBack(roomModel)
			roomModel.RoomState = models.ROOM_STATE_PLAYING_END
			break
		}
	}
}
