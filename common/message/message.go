package message

const (
	LoginMsgType = "LoginMsg"
	LoginResultMsgType = "LoginResultMsg"
	RegisterMsgType = "RegisterMsg"
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMsg struct {
	UserId int `json:"userId"`
	UserPwd string `json:"userPwd"`
	UserName string `json:"userName"`
}

type LoginResultMsg struct {
	Code int `json:"code"`
	Error string `json:"error"`
}