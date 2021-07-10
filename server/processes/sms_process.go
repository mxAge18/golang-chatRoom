package processes

import (
	"encoding/json"
	"fmt"
	"chatPro/common/message"
	"chatPro/server/utils"
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
// 优化发送消息，发送的消息可以点对点，from userid1 to userid2 ,message
// 如果在线直接发送 如果不在线保存到redis
func (this *SmsServerProcess) SendMsgToSomeOne(msg *message.Message) {
	var returnMsg message.Message
	returnMsg.Type = message.SmsMsgSingleReturnType
	var smsSingleMsg message.SmsMsgSingle
	json.Unmarshal([]byte(msg.Data), &smsSingleMsg)
	data, err := json.Marshal(smsSingleMsg)
	if err != nil {
		fmt.Println("smsMsg json.Marshal err", err)
	}
	returnMsg.Data = string(data)
	data, err = json.Marshal(returnMsg)
	if err != nil {
		fmt.Println("msg json.Marshal err", err)
	}
	val, ok := ServerUserManger.onlineUsers[smsSingleMsg.To.UserId]
	if ok {
		this.SendMsgToEachOnlineUser(data, val.Conn)
	} else {
		fmt.Println("用户不在线，存储到服务器")
		// todo
		
	}
}