module.exports = require("protobufjs").newBuilder({})['import']({
    "package": "beans",
    "messages": [
        {
            "name": "CreateRoomBean",
            "fields": [
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "game_id",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "double",
                    "name": "longitude",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "double",
                    "name": "latitude",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "device_info",
                    "id": 5
                }
            ]
        },
        {
            "name": "JoinRoomBean",
            "fields": [
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "game_id",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "player_tocken",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "double",
                    "name": "longitude",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "double",
                    "name": "latitude",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "ip",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "device_info",
                    "id": 6
                }
            ]
        },
        {
            "name": "PlayerDeviceBean",
            "fields": [
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "player_id",
                    "id": 11
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "player_name",
                    "id": 13
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "game_id",
                    "id": 10
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "room_id",
                    "id": 12
                },
                {
                    "rule": "optional",
                    "type": "double",
                    "name": "angle_alpha",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "double",
                    "name": "angle_beta",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "double",
                    "name": "angle_gamma",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "double",
                    "name": "acce_x",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "double",
                    "name": "acce_y",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "double",
                    "name": "acce_z",
                    "id": 6
                },
                {
                    "rule": "optional",
                    "type": "double",
                    "name": "acce_alpha",
                    "id": 7
                },
                {
                    "rule": "optional",
                    "type": "double",
                    "name": "acce_beta",
                    "id": 8
                },
                {
                    "rule": "optional",
                    "type": "double",
                    "name": "acce_gamma",
                    "id": 9
                }
            ]
        },
        {
            "name": "PlagerJoinGameBean",
            "fields": [
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "player_id",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "player_name",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "player_level",
                    "id": 3
                }
            ]
        },
        {
            "name": "PlayerLeaveGameBean",
            "fields": [
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "player_id",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "player_name",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "player_level",
                    "id": 3
                }
            ]
        },
        {
            "name": "HorseSpeedBean",
            "fields": [
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "player_id",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "player_name",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "player_level",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "double",
                    "name": "player_speed",
                    "id": 4
                }
            ]
        },
        {
            "name": "ServerResponseCreateRoomBean",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "max_count",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "room_id",
                    "id": 2
                }
            ]
        },
        {
            "name": "ServerResponseJoinRoomBean",
            "fields": [
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "room_id",
                    "id": 1
                }
            ]
        },
        {
            "name": "ServerSendBean",
            "fields": [
                {
                    "rule": "optional",
                    "type": "int32",
                    "name": "result_code",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "request_id",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "desc",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "RequestOperationCode",
                    "name": "option_code",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "ServerResponseCreateRoomBean",
                    "name": "response_createroom_bean",
                    "id": 5,
                    "oneof": "Bean"
                },
                {
                    "rule": "optional",
                    "type": "ServerResponseJoinRoomBean",
                    "name": "response_joinroom_bean",
                    "id": 6,
                    "oneof": "Bean"
                },
                {
                    "rule": "optional",
                    "type": "PlagerJoinGameBean",
                    "name": "player_join_bean",
                    "id": 7,
                    "oneof": "Bean"
                },
                {
                    "rule": "optional",
                    "type": "PlayerLeaveGameBean",
                    "name": "player_level_bean",
                    "id": 8,
                    "oneof": "Bean"
                },
                {
                    "rule": "optional",
                    "type": "PlayerDeviceBean",
                    "name": "player_device_bean",
                    "id": 9,
                    "oneof": "Bean"
                },
                {
                    "rule": "optional",
                    "type": "HorseSpeedBean",
                    "name": "player_horsesgame_speed_bean",
                    "id": 10,
                    "oneof": "Bean"
                }
            ],
            "oneofs": {
                "Bean": [
                    5,
                    6,
                    7,
                    8,
                    9,
                    10
                ]
            }
        },
        {
            "name": "ClientRequestBean",
            "fields": [
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "request_id",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "RequestOperationCode",
                    "name": "option_code",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "CreateRoomBean",
                    "name": "createroom_bean",
                    "id": 3,
                    "oneof": "Bean"
                },
                {
                    "rule": "optional",
                    "type": "JoinRoomBean",
                    "name": "joinroom_bean",
                    "id": 4,
                    "oneof": "Bean"
                },
                {
                    "rule": "optional",
                    "type": "PlayerDeviceBean",
                    "name": "playerdevice_bean",
                    "id": 5,
                    "oneof": "Bean"
                }
            ],
            "oneofs": {
                "Bean": [
                    3,
                    4,
                    5
                ]
            }
        }
    ],
    "enums": [
        {
            "name": "RequestOperationCode",
            "values": [
                {
                    "name": "REQUEST_OPERATIONCODE_PLAYERDEVICCEBEAN",
                    "id": 0
                },
                {
                    "name": "REQUEST_OPERATIONCODE_CREATEROOM",
                    "id": 1
                },
                {
                    "name": "REQUEST_OPERATIONCODE_JOINROOM",
                    "id": 2
                }
            ]
        },
        {
            "name": "SendMessageOperationCode",
            "values": [
                {
                    "name": "SENDMESSAGE_OPERATIONCODE_RESPONSE",
                    "id": 0
                },
                {
                    "name": "SENDMESSAGE_OPERATIONCODE_PLAYERJOINGAME",
                    "id": 1
                },
                {
                    "name": "SENDMESSAGE_OPERATIONCODE_PLAYERDEVICE",
                    "id": 2
                },
                {
                    "name": "SENDMESSAGE_OPERATIONCODE_HORSEGAME_SPEED",
                    "id": 3
                }
            ]
        }
    ]
}).build("beans");

///**
// * 将js对象转成protobuf的二进制数据
// * msgName 对应proto里面的消息名称
// * obj是msgName对应的js对象
// **/
//module.exports.encodeObject = function ( msgName, obj )
//{
//    try {
//        var msgObj = new horse[msgName](obj);
//        var buffer = msgObj.encode().toBuffer();
//        return buffer;
//    } catch (e) {
//        console.log(e);
//        return new ArrayBuffer();
//    }
//}
///**
// * 将protobuf的二进制数据 转成js对象
// * msgName 对应proto里面的消息名称
// * buffer
// **/
//module.exports.decodeBuffer = function ( msgName, buffer )
//{
//    try {
//        var message = horse[msgName].decode(buffer)
//        return message;
//    } catch (e) {
//        console.log(e);
//        return {};
//    }
//}


