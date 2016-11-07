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
	AppointmentWaitJoinPlayerChan = make(chan string, 500)              //预约玩家排队等待进入房间Channel(高优先级)
	WaitRPCCheckPlayerChan        = make(chan string, 500)              //等待PRC验证Channel
	RPCCheckSuccessSignChan       = make(chan string, 500)              //rpc验证成功信号
	RPCCheckErrorSignChan         = make(*beans.RPCResponse, 500)       //rpc验证失败，玩家没有权限
	StoragePlayerInstanceChan     = make(chan *models.PlayerModel, 500) //存储PlayerModel实体Channel
	WaitJoinRommPlayerChan        = make(*models.PlayerModel, 500)      //玩家排队等待进入房间Channel
)
var (
	PlayerId2PlayerInstanceMap = smap.New() //记录所有的用户实体
)

func init() {
	go loop_storage_playerInstance()
	go loop_rpc_check()
	go loop_player_property_dispatch()
}

func loop_storage_playerInstance() {
	for {
		select {
		case playerModel := <-StoragePlayerInstanceChan:
			PlayerId2PlayerInstanceMap.Insert(playerModel.PlayerId, playerModel)
			WaitRPCCheckPlayerChan <- playerModel.PlayerId
		}
	}
}

/**
 * RPC玩家验证模块
 * @return {[type]} [description]
 */
func loop_rpc_check() {
	for {
		select {
		case playerID := <-WaitRPCCheckPlayerChan:
			playerModel := (PlayerId2PlayerInstanceMap.Get(playerID)).(*models.PlayerModel)
			go checkPlayer(playerModel)
		case rpcResponseBean := <-RPCCheckErrorSignChan: //rpc验证失败
		/**
		 * 错误处理,返回给客户端提示稍后重试
		 */
		case playerID := <-RPCCheckSuccessSignChan: //rpc验证成功
			WaitJoinRommPlayerChan <- playerID
		}
	}
}

/**
 * 从BAZIRIM验证tocken，并且获取用户名和等级
 * @param  {[type]} player *models.Player [description]
 * @return {[type]}        [description]
 */
func checkPlayer(player *models.PlayerModel) {
	url := beego.AppConfig.String("rpc::checkplayerurl")
	req := httplib.Post(url)
	responseBean := beans.RPCResponse{}
	err := req.ToJSON(&responseBean)
	if err != nil {
		responseBean.ErrorCode = beego.AppConfig.String("errcode::rpc_response_format_error")
		responseBean.Desc = "内部错误:远程服务器返回信息格式有误"
		RPCCheckErrorSignChan <- responseBean
		return
	}
	if responseBean.ErrorCode != "0" { //验证失败
		RPCCheckErrorSignChan <- responseBean
		return
	}
	checkBean, err := responseBean.Data.(*beans.UserCheckData)
	if err != nil {
		responseBean.ErrorCode = beego.AppConfig.String("errcode::rpc_response_format_error")
		responseBean.Desc = "内部错误:远程服务器返回信息格式有误"
		RPCCheckErrorSignChan <- responseBean
		return
	}
	player.PlayerLevel = int32(checkBean.Level)
	player.PlayerName = checkBean.Name
	player.UserId = checkBean.UserId
	RPCCheckSuccessSignChan <- player.PlayerId
}

/**
 * 玩家属性分配模块,查看Redies判断是否是预约玩家并进行相应的分组
 * @return {[type]} [description]
 */
func loop_player_property_dispatch() {
	for {
		select {
		case playerID := <-WaitJoinRommPlayerChan:
			playerModel := (PlayerId2PlayerInstanceMap.Get(playerID)).(*models.RoomModel) //此刻是在rpc验证之后，所以playerModel.UserId有效,需要在Redies中查找相应的预约ID(AppointmentId)

		}
	}
}

func PlayerJoin(requestId string, conn *websocket.Conn, bean *beans.JoinRoomBean) {
	player := models.NewPlayer(conn, bean.PlayerTocken, bean.Longitude, bean.Latitude, bean.DeviceInfo)
	StoragePlayerInstanceChan <- player
}
