package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"hiphopkys/servers/horse/beans"
	"hiphopkys/servers/horse/caches"
	"hiphopkys/servers/horse/models"
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
	if nil != dataJson {
		responseBean.ErrorCode = beego.AppConfig.String("errcode::rpc_request_setappointment_error_code")
		responseBean.Desc = beego.AppConfig.String("errcode::rpc_request_setappointment_error_desc")
	}
	buffer, err := json.Marshal(&responseBean)
	if err != nil {
		beego.BeeLogger.Error("PushAppointmentUser::json.Marshal(responseBean) error：%s", err.Error())
		this.Ctx.WriteString(err.Error())
		return
	}
	appointmentUserModel := beans.AppointmentUser{}
	err = json.Unmarshal([]byte(dataJson), &appointmentUserModel)
	if err != nil {
		responseBean.ErrorCode = beego.AppConfig.String("errcode::rpc_request_setappointment_error_code")
		responseBean.Desc = beego.AppConfig.String("errcode::rpc_request_setappointment_error_desc" + ",err:" + err.Error())
	} else {

	}
	jsonString := string(buffer)
	this.Ctx.WriteString(jsonString)
}
