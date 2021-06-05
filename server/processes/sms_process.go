package processes

import (
	"encoding/json"
	"fmt"
	"go_code/chatPro/common/message"
	"go_code/chatPro/server/utils"
	"net"
)

type SmsServerProcess struct{

}

func (this *SmsServerProcess) SendGroupMsg(msg *message.Message) {
	var returnMsg message.Message
	returnMsg.Type = message.GroupReturnMsgType
	var smsMsg message.SmsMsg
	json.Unmarshal([]byte(msg.Data), &smsMsg)
	data, err := json.Marshal(smsMsg)
	if err != nil {
		fmt.Println("smsMsg json.Marshal err", err)
	}
	returnMsg.Data = string(data)
	data, err = json.Marshal(returnMsg)
	if err != nil {
		fmt.Println("msg json.Marshal err", err)
	}
	for id, up := range(ServerUserManger.onlineUsers) {
		if id == smsMsg.User.UserId {
			continue
		}
		this.SendMsgToEachOnlineUser(data, up.Conn)
	}
}

func (this *SmsServerProcess) SendMsgToEachOnlineUser(data []byte, conn net.Conn) {
	tr := &utils.Transfer{
		Conn: conn,
	}
	err := tr.WritePkg(data)
	if err != nil {
		fmt.Println("error of Sms send single msg", err)
	}
	return
}