package processes

import (
	"encoding/json"
	"fmt"
	"chatPro/common/message"
	"chatPro/server/utils"
)

type SmsProcess struct {

}

// 发送群聊消息
func (this *SmsProcess) SendGroup(content string) (err error) {
	var msg message.Message
	msg.Type = message.SmsMsgType
	var smsMsg message.SmsMsg
	smsMsg.Body = content
	smsMsg.User = CurrentUserObj.User
	data, err := json.Marshal(smsMsg)
	if err != nil {
		fmt.Println("smsMsg json.Marshal err", err)
		return
	}
	msg.Data = string(data)
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("msg json.Marshal err", err)
		return
	}
	tr := &utils.Transfer{
		Conn: CurrentUserObj.Conn,
	}
	err = tr.WritePkg(data)
	if err != nil {
		fmt.Println("sms WritePkg json.Marshal err", err)
		return
	}
	return
}

func (this *SmsProcess) ReadGroupMsg(msg message.Message) (err error) {
	var groupMsg message.SmsMsg
	json.Unmarshal([]byte(msg.Data), &groupMsg)
	fmt.Println("发信人：", groupMsg.User.UserId)
	fmt.Println("发信内容：", groupMsg.Body)
	return
}