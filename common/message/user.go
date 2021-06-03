package message

type User struct {
	UserId string `json:"userId"`
	UserName string `json:"userName"`
	UserPwd string `json:"userPwd"`
	UserStatus int `json:"userStatus"` // user online departure not-online
}

