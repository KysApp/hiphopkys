package models

import (
	"github.com/gorilla/websocket"
	"labix.org/v2/mgo/bson"
	"strings"
)

type PlayerStateValue int32

const (
	PLAYER_STATE_CHECKING      PlayerStateValue = iota //正在RPC验证用户
	PLAYER_STATE_WAITJOIN                              //没有可用房间,排队中
	PLAYER_STATE_WAITGAMESTART                         //已加入房间,但房间未满,等待房间满员开始游戏
	PLAYER_STATE_PLAYING                               //正在游戏中
	PLAYER_STATE_END                                   //一轮游戏结束,但玩家还在房间中
)

type PlayerModel struct {
	PlayerId       string
	Tocken         string           //客户端Tocken，需要通过RPC从BAZIRIM验证玩家是否有权限游戏
	RoomId         string           //加入到的房间的ID
	PlayerState    PlayerStateValue //玩家当前状态
	Longitude      float64          //PC端经度,可为空
	Latitude       float64          //PC端纬度,可为空
	DeviceInfo     string           //PC端描述,可为空
	Conn           *websocket.Conn  //PC端WebSocket链接
	Ip             string           //客户端IP
	NetworkSegment string           //客户端网段
	AppointmentId  string           //预约ID
	/**
	 *以下数据是通过RPC从BAZIRIM获得
	 */
	PlayerName  string //玩家昵称
	PlayerLevel int32  //玩家等级
	UserId      string //玩家ID(BAZIRIM中对应的)
}

func NewPlayerModel(conn *websocket.Conn, tocken string, longitude float64, latitude float64, deviceinfo string) *PlayerModel {
	player := &PlayerModel{
		PlayerId:    bson.NewObjectId().Hex(),
		Conn:        conn,
		Tocken:      tocken,
		PlayerState: PLAYER_STATE_CHECKING,
		Longitude:   longitude,
		Latitude:    latitude,
		DeviceInfo:  deviceinfo,
		Ip:          conn.RemoteAddr().String(),
	}
	index := strings.LastIndex(player.Ip, ".")
	player.NetworkSegment = (player.Ip)[0:index]
	return player
}
