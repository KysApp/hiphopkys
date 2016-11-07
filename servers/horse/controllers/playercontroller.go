package controllers

import (
	"container/list"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/gorilla/websocket"
	"hiphopkys/servers/commons/container/list"
	"hiphopkys/servers/commons/container/smap"
	"hiphopkys/servers/horse/beans"
	"hiphopkys/servers/horse/models"
	"sync"
)

var (
	AppointmentWaitJoinPlayerChan = make(chan string, 500) //预约玩家排队等待进入房间Channel(高优先级)
)
var (
	PlayerId2RoomInstanceMap = smap.New() //记录所有的用户实体
)
var (
	WaitCheckPlayerSafeList = struct { //玩家刚刚加入后加入到此队列,表示等待PRC验证Tocken
		sync.RWMutex
		List *list.List
	}{List: list.New()}

	WaitJoinRommPlayerList = struct { //玩家PRC验证完tocken后加入此队列等待加入房间
		sync.RWMutex
		List *list.List
	}{List: list.New()}

	JoinRommWaitPlayingPlayerSafeMap = struct { //加入房间以后
		sync.RWMutex
		Map map[string]*models.Player
	}{Map: make([string]*models.Player)}
)

var (
	WaitJoinRommPlayerChan = make(*models.Player, 500)
)

func init() {
	go loop()
}

func loop() {
	for {
		select {
		case player := <-WaitJoinRommPlayerChan:
			// WaitJoinRommPlayerList.PushBack(player)
		}
	}
}

/**
 * 从BAZIRIM验证tocken，并且获取用户名和等级
 * @param  {[type]} player *models.Player [description]
 * @return {[type]}        [description]
 */
func checkPlayer(player *models.Player) {
	url := beego.AppConfig.String("rpc::checkplayerurl")
	req := httplib.Post(url)
	responseBean := beans.RPCResponse{}
	err := req.ToJSON(&responseBean)
	if err != nil {
		beego.BeeLogger.Error("RPC::checkPlayer错误,返回结果不合法:%s", err.Error())
		/**
		 * 错误处理,返回给客户端提示稍后重试
		 */
		// player.Conn.Close()
		return
	}
	if responseBean.ErrorCode != "0" { //不成功
		beego.BeeLogger.Error("RPC::checkPlayer错误:%s", responseBean.Desc)
		/**
		 * 错误处理,用户没有权限或其他错误,返回个客户端
		 */
		return
	}
	checkBean, err := responseBean.Data.(*beans.UserCheckData)
	if err != nil {
		beego.BeeLogger.Error("RPC::checkPlayer错误,返回结果不合法:%s", err.Error())
		/**
		 * 错误处理,返回给客户端提示稍后重试
		 */
		// player.Conn.Close()
		return
	}
	player.PlayerLevel = int32(checkBean.Level)
	player.PlayerName = checkBean.Name
	player.UserId = checkBean.UserId
}

func PlayerJoin(requestId string, conn *websocket.Conn, bean *beans.JoinRoomBean) {
	player := models.NewPlayer(conn, bean.PlayerTocken, bean.Longitude, bean.Latitude, bean.DeviceInfo)
	go checkPlayer(player)
}
