package models

import (
	"container/list"
	"github.com/gorilla/websocket"
	"labix.org/v2/mgo/bson"
)

type RoomStateValue int32

const (
	ROOM_STATE_WAITJOIN          RoomStateValue = iota //未满,正在等待玩家加入的房间队列
	ROOM_STATE_FULL_WAIT_PLAYING                       //已满,等待游戏开始的房间队列
	ROOM_STATE_PLAYING                                 //已满,游戏正在进行中的房间队列
)

type RoomModel struct {
	RoomId           string
	GameId           string          //游戏Id,与CreateRoomBean中的game_id相同,由PC传入
	Longitude        float64         //PC端经度,可为空
	Latitude         float64         //PC端纬度,可为空
	DeviceInfo       string          //PC端描述,可为空
	RoomState        RoomStateValue  //房间状态
	Conn             *websocket.Conn //PC端WebSocket链接
	PlayerTockenList *list.List      //房间中玩家tocken队列
}

func NewRoomModel(conn *websocket.Conn, gameId string, longitude float64, latitude float64, deviceInfo string) *RoomModel {
	room := &RoomModel{
		RoomId:           bson.NewObjectId().Hex(),
		GameId:           gameId,
		Longitude:        longitude,
		Latitude:         latitude,
		RoomState:        ROOM_STATE_WAITJOIN,
		Conn:             conn,
		PlayerTockenList: list.New(),
	}
	return room
}
