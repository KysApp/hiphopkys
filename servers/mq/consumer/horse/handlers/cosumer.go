package handlers

import (
	"fmt"
	"hiphopkys/servers/mq/consumer/horse/beans"
)

func PrintSelf() {
	fmt.Println("这是horse_consumer V1.0")
}

func createBean() beans.User {
	return beans.User{
		Name: "test-user",
		Id:   "test-id",
		Age:  20,
		Con: beans.Contact{
			Tel:   "13333333333",
			Email: "abc@gmail.com",
		},
	}
}

func init() {
	userBean := createBean()
	modelBuf, err := userBean.Marshal(nil)
	if err != nil {
		fmt.Printf("Marshal错误:%s\n", err.Error())
		return
	}
	fmt.Printf("编码UserBean:%#v成功，buf:%v\n", userBean, modelBuf)
	readBean := beans.User{}
	_, err = readBean.Unmarshal(modelBuf)
	if nil == err {
		fmt.Printf("解码成功:%#v\n", readBean)
	}
	PrintSelf()
}
