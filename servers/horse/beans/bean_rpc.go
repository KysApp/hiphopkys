package beans

type UserCheckData struct {
	Name     string `json:"name"`
	Level    int32  `json:"level"`
	UserId   string `json:"user_id"`
	PlayerId string `json:"-"`
}
type AppointmentUser struct {
	UserId             string `json:"user_id"`
	AppointmentId      string `json:"appointment_id"`
	LimiteCanPlayCount int32  `json:"can_play_count"`
}

type RPCResponse struct {
	ErrorCode string      `json:"errorCode"`
	Desc      string      `json:"desc"`
	Data      interface{} `json:"data"`
}
