//var Beans = require("beans_horse");
var ByteBuffer = require("bytebuffer");
var ProtoBuf = require("protobufjs");
var Models = require("models");
var Message = require("Message").Message;
cc.Class({
    extends: cc.Component,

    properties: {
        label: {
            default: null,
            type: cc.Label
        },
        // defaults, set visually when attaching this script to the Canvas
        text: 'Hello, World!',
        conn: null,
        counter: 0,
        builder: null,
        beans: null,
        gameInfoModel: null,
    },

    //对玩家加入房间后响应的处理
    wsProcess_PlayerJoinResponse: function(event){
        event.stopPropagation(); //不向下传递
        var joinRespBean = event.getUserData();
        //process
    },


    // use this for initialization
    onLoad: function () {
        this.label.string = this.text;
        var self = this;
        var protoFile = "protobuf/beans_horse";
        this.gameInfoModel = new Models.GameInfoModel();
        this.gameInfoModel._playerTocken = "tocken0";
        this.gameInfoModel._gameId = "horse-v1.0";
        this.gameInfoModel._deviceInfo = "测试手机客户端设备0";
        this.gameInfoModel._latitude = 99.99999;
        this.gameInfoModel._longitude = 55.55555;
        //注册事件处理
        this.node.on(Message.MESSAGE_WS_JOINRESPONSE_PROCESS,this.wsProcess_PlayerJoinResponse,this);
        cc.loader.loadRes(protoFile,function(err, bgTexture2D){
            if(err){
                console.log("加载资源错误"+JSON.stringify(err));
                return;
            }
            self.builder = ProtoBuf.protoFromString(bgTexture2D);
            self.beans = self.builder.build("beans");
            self.conn = new WebSocket("ws://127.0.0.1:5052/game/join");
            self.conn.onopen = function (event) {
                console.log("连接成功");
                var joinRoomBean = new self.beans.JoinRoomBean();
                joinRoomBean.device_info = self.gameInfoModel._deviceInfo ;
                joinRoomBean.game_id = self.gameInfoModel._gameId;
                joinRoomBean.latitude = self.gameInfoModel._latitude;
                joinRoomBean.longitude = self.gameInfoModel._longitude;
                joinRoomBean.player_tocken = self.gameInfoModel._playerTocken;
                var requestBean = new self.beans.ClientRequestBean();
                requestBean.option_code = self.beans.RequestOperationCode.REQUEST_OPERATIONCODE_JOINROOM;
                requestBean.request_id = "request-0";
                requestBean.joinroom_bean = joinRoomBean;
                var msgBuf = requestBean.encode().toArrayBuffer();
                self.conn.send(msgBuf);
            };
            self.conn.onerror = function (event) {
                console.log("Send Text fired an error");
            };
            self.conn.onclose = function (event) {
                console.log("WebSocket instance closed.");
            };
            self.conn.onmessage = function (event) {
                console.log("response text msg: " + event.data);

                //var serverSendBean = new self.beans.ServerSendBean();
                var msg = self.beans.ServerSendBean.decode(event.data);
                var joinrespBean = event.data;

                /**
                 * 如果是joingame的响应
                 * */
                var event_joinresp = new cc.Event.EventCustom(Message.MESSAGE_WS_JOINRESPONSE_PROCESS, true);
                event_joinresp.setUserData(joinrespBean)
                self.node.dispatchEvent(event_joinresp);

            };

        });



        //setTimeout(function () {
        //    if (self.conn.readyState === WebSocket.OPEN) {
        //        this.conn.send("Hello WebSocket, I'm a text message.");
        //        var protoFile = "protobuf/beans_horse";
        //        cc.loader.loadRes(protoFile, function (err, bgTexture2D){
        //            cc.log("loadfinish");
        //            self.builder = ProtoBuf.protoFromString(bgTexture2D);
        //            self.beans = self.builder.build("beans");
        //
        //
        //            for(var i = 2 ; i != 5; ++i){
        //                var joinRoomBean = new self.beans.JoinRoomBean();
        //                joinRoomBean.device_info = "测试手机客户端设备" + i;
        //                joinRoomBean.game_id = "horse-v1.0";
        //                joinRoomBean.latitude = 100.5 + i;
        //                joinRoomBean.longitude = 100.99 + i;
        //                joinRoomBean.player_tocken = "tocken" + i;
        //                var requestBean = new self.beans.ClientRequestBean();
        //                requestBean.option_code = self.beans.RequestOperationCode.REQUEST_OPERATIONCODE_JOINROOM;
        //                requestBean.request_id = "request-0"+i;
        //                requestBean.joinroom_bean = joinRoomBean;
        //                var msgBuf = requestBean.encode().toArrayBuffer();
        //                self.conn.send(msgBuf);
        //            }
        //
        //
        //        });
        //    }
        //    else {
        //        console.log("WebSocket instance wasn't ready...");
        //    }
        //}, 10);








    },



    // called every frame
    update: function (dt) {

    },
});
