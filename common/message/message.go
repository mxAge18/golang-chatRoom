package message

const (
	LoginMsgType            = "LoginMsg"
	LoginMesType            = "LoginMes"
	LoginResultMsgType      = "LoginResultMsg"
	RegisterMsgType         = "RegisterMsg"
	NotifyUserStatusMsgType = "NotifyUserStatusMsg"
)

const (
	UserOnline     = 1
	UserOffline    = 2
	UserBusyStatus = 3
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMsg struct {
	UserId   string `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

type LoginResultMsg struct {
	Code  int      `json:"code"`
	Error string   `json:"error"`
	Data  []string `json:"data"`
}

type RegisterMsg struct {
	User User `json:"user"`
}

type RegisterResultMsg struct {
	Code  int    `json:"code"` // 400标识已占用  200 标识ok // 403标识User未校验通过
	Error string `json:"error"`
	// Data string `json:"data"`
}

// user online information of status
type NotifyUserStatusMsg struct {
	UserId     string `json:"userId"`
	UserName   string `json:"userName"`
	UserStatus int    `json:"userStatus"`
}
