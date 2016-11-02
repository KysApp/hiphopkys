package beans

type UserCheckData struct {
	Name  string `json:"name"`
	Level int32  `json:"level"`
}

type RPCResponse struct {
	ErrorCode string      `json:"errorCode"`
	Desc      string      `json:"desc"`
	Data      interface{} `json:"data"`
}
