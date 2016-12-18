package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/gorilla/websocket"
	// "hiphopkys/servers/commons/container/list"
	"hiphopkys/servers/commons/container/smap"
	"hiphopkys/servers/horse/beans"
	"hiphopkys/servers/horse/caches"
	"hiphopkys/servers/horse/models"
	// "sync"
)

var (
	AppointmentWaitJoinPlayerChan = make(chan string, 500)              //预约玩家排队等待进入房间Channel(高优先级)
	WaitRPCCheckPlayerChan        = make(chan string, 500)              //等待PRC验证Channel
	RPCCheckSuccessSignChan       = make(chan string, 500)              //rpc验证成功信号
	RPCCheckErrorSignChan         = make(chan *beans.RPCResponse, 500)  //rpc验证失败，玩家没有权限
	StoragePlayerInstanceChan     = make(chan *models.PlayerModel, 500) //存储PlayerModel实体Channel
	WaitJoinRommPlayerChan        = make(chan string, 500)              //玩家排队等待进入房间Channel
	OnePlayerJoinSign             = make(chan string, 500)
)
var (
	PlayerId2PlayerInstanceMap = smap.New() //记录所有的用户实体
)

func init() {
	go loop_storage_playerInstance()
	go loop_rpc_check()
	go loop_player_property_dispatch()
	go loop_onePlayerJoinedRoom()
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

func loop_onePlayerJoinedRoom() {
	for {
		select {
		case playerId := <-OnePlayerJoinSign:
			playerModel := (PlayerId2PlayerInstanceMap.Get(playerId)).(*models.PlayerModel)
			/**
			 * 通知TV端新加入玩家信息
			 */
			wsPlayerJoinBean := &beans.ServerSendBean_PlayerJoinBean{
				PlayerJoinBean: &beans.PlagerJoinGameBean{},
			}
			wsPlayerJoinBean.PlayerJoinBean.PlayerId = playerId
			wsPlayerJoinBean.PlayerJoinBean.PlayerLevel = playerModel.PlayerLevel
			wsPlayerJoinBean.PlayerJoinBean.PlayerName = playerModel.PlayerName
			wsPlayerJoinBean.PlayerJoinBean.PlayerTocken = playerModel.Tocken
			wsResponseBean := beans.ServerSendBean{}
			wsResponseBean.Bean = wsPlayerJoinBean
			wsResponseBean.OptionCode = beans.SendMessageOperationCode_SENDMESSAGE_OPERATIONCODE_PLAYERJOINGAME
			wsResponseBean.ResultCode = "0"
			wsResponseBean.Desc = "有新玩家加入"
			if msgBuf, err := wsResponseBean.Marshal(); nil == err {
				roomModel := (RoomId2RoomInstanceMap.Get(playerModel.RoomId)).(*models.RoomModel)
				if err := roomModel.Conn.WriteMessage(websocket.BinaryMessage, msgBuf); err != nil {
					beego.BeeLogger.Error("通知TV端新玩家加入时websocket错误:%s", err.Error())
				}

			} else {
				/**
				 *
				 */
				beego.BeeLogger.Error("通知TV端新玩家加入时protobuf错误:%s", err.Error())

			}

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
			 * 错误处理,返回给客户端错误提示
			 */
			beego.BeeLogger.Error("rpc验证失败:%#v", rpcResponseBean)
			userCheckData, ok := rpcResponseBean.Data.(*beans.UserCheckData)
			wsResponseBean := beans.ServerSendBean{}
			wsResponseBean.ResultCode = rpcResponseBean.ErrorCode
			wsResponseBean.Desc = rpcResponseBean.Desc
			wsResponseBean.OptionCode = beans.SendMessageOperationCode_SENDMESSAGE_OPERATIONCODE_RESPONSE_PLAYERJOINBEAN
			if ok {
				playerModel := (PlayerId2PlayerInstanceMap.Get(userCheckData.PlayerId)).(*models.PlayerModel)
				msgBuf, err := wsResponseBean.Marshal()
				if err != nil {
					beego.BeeLogger.Error("内部错误:玩家验证失败，当给客户端返回错误信息是protobuf解析失败:%s,强行关闭客户端连接", err.Error())
					playerModel.Conn.Close()
				} else {
					playerModel.Conn.WriteMessage(websocket.BinaryMessage, msgBuf)
				}
			} else {
				beego.BeeLogger.Error("内部错误:玩家验证失败，但是无法查找到playerId")
			}
		case playerID := <-RPCCheckSuccessSignChan: //rpc验证成功
			WaitJoinRommPlayerChan <- playerID
			/**
			 * 此时暂不用给客户端返回,在loop_player_property_dispatch()中判断完是否有房间后再做处理
			 */

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
	req.Param("tocken", player.Tocken)
	responseBean := &beans.RPCResponse{}
	userCheckData := &beans.UserCheckData{}
	userCheckData.PlayerId = player.PlayerId
	responseBean.Data = userCheckData
	err := req.ToJSON(responseBean)
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
	checkBean, isOK := responseBean.Data.(*beans.UserCheckData)
	if !isOK {
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
			playerModel := (PlayerId2PlayerInstanceMap.Get(playerID)).(*models.PlayerModel) //此刻是在rpc验证之后，所以playerModel.UserId有效,需要在Redies中查找相应的预约ID(AppointmentId)
			isAppointment, appointmentModel := caches.CachePullAppointmentUser(playerModel.UserId)
			if isAppointment {
				playerModel.AppointmentId = appointmentModel.AppointmentId
				isHas := false
				WaitPlayerJoinRoomList.Map(func(rommID_Value interface{}) bool {
					roomID := rommID_Value.(string)
					roomModel := (RoomId2RoomInstanceMap.Get(roomID)).(*models.RoomModel)
					if playerModel.AppointmentId == roomModel.AppointmentId {
						roomModel.PlayerUserIdList.PushBack(playerID)
						playerModel.RoomId = roomID
						OnePlayerJoinSign <- playerModel.PlayerId
						AppointmentWaitCheckStateRoomChan <- roomID
						isHas = true
						beego.BeeLogger.Error("受邀玩家信息(此时已存在房间):%#v", playerModel)
						/**
						 *
						 */
						wsResponseBean := beans.ServerSendBean{}
						wsResponseBean.ResultCode = beego.AppConfig.String("errcode::check_player_waitplay_error_code")
						wsResponseBean.Desc = beego.AppConfig.String("errcode::check_player_waitplay_error_desc")
						wsResponseBean.OptionCode = beans.SendMessageOperationCode_SENDMESSAGE_OPERATIONCODE_RESPONSE_PLAYERJOINBEAN
						wsResponseBean.Bean = &beans.ServerSendBean_ResponseJoinroomBean{
							ResponseJoinroomBean: &beans.ServerResponseJoinRoomBean{
								RoomId:   playerModel.RoomId,
								PlayerId: playerModel.PlayerId,
							},
						}

						playerModel := (PlayerId2PlayerInstanceMap.Get(playerModel.PlayerId)).(*models.PlayerModel)
						if msgBuf, err := wsResponseBean.Marshal(); err == nil {
							playerModel.Conn.WriteMessage(websocket.BinaryMessage, msgBuf)
						} else {
							beego.BeeLogger.Error("加入房间成功，但给客户端反馈时的response protobuf marshal失败:%s,这里强制关闭连接", err.Error())
							playerModel.Conn.Close()
						}

						return false
					}
					return true
				})
				if !isHas {
					beego.BeeLogger.Error("受邀玩家信息(此时还不存在房间):%#v", playerModel)
					wsResponseBean := beans.ServerSendBean{}
					wsResponseBean.ResultCode = beego.AppConfig.String("errcode::check_player_waitjoin_error_code")
					wsResponseBean.Desc = beego.AppConfig.String("errcode::check_player_waitjoin_error_desc")
					wsResponseBean.OptionCode = beans.SendMessageOperationCode_SENDMESSAGE_OPERATIONCODE_RESPONSE_PLAYERJOINBEAN
					wsResponseBean.Bean = &beans.ServerSendBean_ResponseJoinroomBean{
						ResponseJoinroomBean: &beans.ServerResponseJoinRoomBean{
							PlayerId: playerModel.PlayerId,
						},
					}
					playerModel := (PlayerId2PlayerInstanceMap.Get(playerID)).(*models.PlayerModel)
					if msgBuf, err := wsResponseBean.Marshal(); err == nil {
						playerModel.Conn.WriteMessage(websocket.BinaryMessage, msgBuf)

					} else {
						beego.BeeLogger.Error("验证成功,且等待加入房间，但给客户端反馈时的response protobuf marshal失败:%s,这里强制关闭连接", err.Error())
						playerModel.Conn.Close()
					}
					AppointmentWaitJoinPlayerChan <- playerID
				}
			} else {

			}

		}
	}
}

func PlayerJoin(requestId string, conn *websocket.Conn, bean *beans.JoinRoomBean) {
	player := models.NewPlayerModel(conn, bean.PlayerTocken, bean.Longitude, bean.Latitude, bean.DeviceInfo)
	Conn2PlayerIdSmap.Insert(player.Conn, player.PlayerId)
	StoragePlayerInstanceChan <- player
}
