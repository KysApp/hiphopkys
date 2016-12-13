package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"hiphopkys/servers/horse/beans"
	"hiphopkys/servers/horse/caches"
	"hiphopkys/servers/horse/models"
	"time"
)

type RpcController struct {
	beego.Controller
}

/**
 * 添加受邀玩家
 * @param  {[type]} this *RpcController) PushAppointmentUser( [description]
 * @return {[type]}      [description]
 */
func (this *RpcController) PushAppointmentUser() {
	dataJson := this.GetString("data")
	responseBean := beans.RPCResponse{}
	if "" == dataJson {
		responseBean.ErrorCode = beego.AppConfig.String("errcode::rpc_request_setappointment_error_code")
		responseBean.Desc = beego.AppConfig.String("errcode::rpc_request_setappointment_error_desc")
	}
	appointmentUserBean := beans.AppointmentUser{}
	err := json.Unmarshal([]byte(dataJson), &appointmentUserBean)
	if err != nil {
		responseBean.ErrorCode = beego.AppConfig.String("errcode::rpc_request_setappointment_error_code")
		responseBean.Desc = beego.AppConfig.String("errcode::rpc_request_setappointment_error_desc" + ",err:" + err.Error())
	} else {
		responseBean.ErrorCode = "0"
		responseBean.Desc = "success"
		model := &models.AppointmentPlayerCacheModel{}
		model.UserId = appointmentUserBean.UserId
		model.LimiteCanPlayCount = appointmentUserBean.LimiteCanPlayCount
		model.AppointmentId = appointmentUserBean.AppointmentId
		model.AppointmentTimestamp = time.Now().Unix()
		errcode, desc := caches.CachePushAppointmentUser(model)
		responseBean.ErrorCode = errcode
		responseBean.Desc = desc
	}
	buffer, err := json.Marshal(&responseBean)
	if err != nil {
		beego.BeeLogger.Error("PushAppointmentUser:%s", err.Error())
	}
	jsonString := string(buffer)
	this.Ctx.WriteString(jsonString)
}
