package processes

import (
	"encoding/json"
	"fmt"
	"go_code/chatPro/common/message"
	"go_code/chatPro/server/model"
	"go_code/chatPro/server/utils"
	"net"
)


type UserProcessor struct {
	Conn net.Conn

}
func (this *UserProcessor) ServerProcessLogin(msg *message.Message) (err error) {
	var loginMsg message.LoginMsg
	err = json.Unmarshal([]byte(msg.Data), &loginMsg)
	if err != nil {
		return
	}
	var resMsg message.Message
	resMsg.Type = message.LoginResultMsgType
	var loginResMsg message.LoginResultMsg

	user, err := model.MyUserDBO.Login(loginMsg.UserName, loginMsg.UserPwd)

	if err != nil {
		if err == model.ERROR_USER_NOT_EXIST {
			loginResMsg.Code = 500
			loginResMsg.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMsg.Code = 403
			loginResMsg.Error = err.Error()
		} else {
			loginResMsg.Code = 505
			loginResMsg.Error = "服务器内部错误"
		}
	} else {
		loginResMsg.Code = 200
		fmt.Println("登录成功", user)
	}
	data, err := json.Marshal(loginResMsg)
	if err != nil {
		fmt.Println("Json Marshal error", err)
		return
	}
	resMsg.Data = string(data)

	data, err = json.Marshal(resMsg)
	if err != nil {
		fmt.Println("resMsg Json Marshal error", err)
		return
	}
	// data is the login result needs to send to the client, need coding the logic of sending data

	tr := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tr.WritePkg(data)
	if err != nil {
		fmt.Println("writePkg(conn, data) error", err)
		return
	}
	return
}
func (this *UserProcessor) ServerProcessRegister(msg *message.Message) (err error) {
	var registerMsg message.RegisterMsg
	err = json.Unmarshal([]byte(msg.Data), &registerMsg)
	if err != nil {
		return
	}
	var resMsg message.Message
	resMsg.Type = message.RegisterMsgType
	var registerResultMsg message.RegisterResultMsg

	err = model.MyUserDBO.Register(&registerMsg.User)
	if err != nil {
		if err == model.ERROR_USER_EXIST {
			registerResultMsg.Code = 403
			registerResultMsg.Error = model.ERROR_USER_EXIST.Error()
		} else {
			registerResultMsg.Code = 500
		}

	} else {
		registerResultMsg.Code = 200
	}
	data, err := json.Marshal(registerResultMsg)
	if err != nil {
		fmt.Println("json.Marshal(registerResultMsg) error", err)
		return
	}
	resMsg.Data = string(data)

	data, err = json.Marshal(resMsg)
	if err != nil {
		fmt.Println("resMsg Json Marshal error", err)
		return
	}

	tr := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tr.WritePkg(data)
	if err != nil {
		fmt.Println("writePkg(conn, data) error", err)
		return
	}
	return
}
