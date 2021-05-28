package processes

import (
	"encoding/json"
	"fmt"
	"go_code/chatPro/common/message"
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

	if loginMsg.UserName == "maxu" && loginMsg.UserPwd == "123456" {
		loginResMsg.Code = 200
		fmt.Println("登录成功")
	} else {
		loginResMsg.Code = 500
		loginResMsg.Error = "用户名不存在"
		fmt.Println("用户名不存在")
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