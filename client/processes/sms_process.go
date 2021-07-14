package processes

import (
	"chatPro/common/message"
	"chatPro/server/utils"
	"encoding/json"
	"fmt"
	"strings"
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

// 发送消息给某个用户
func (this *SmsProcess) SendSingleMsg(content string, toId string) (err error) {
	var msg message.Message
	msg.Type = message.SmsMsgSingleType
	var smsMsg message.SmsMsgSingle
	smsMsg.Body = content
	smsMsg.From = CurrentUserObj.User
	smsMsg.To = toId
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

func (this *SmsProcess) ReadSingleMsg(msg message.Message) (err error) {
	var singleMsg message.SmsMsg
	json.Unmarshal([]byte(msg.Data), &singleMsg)
	fmt.Println("私信发信人：", singleMsg.User.UserId)
	fmt.Println("发信内容：", singleMsg.Body)
	return
}
func (this *SmsProcess) GetUnreadMsg() (err error) {
	var msg message.Message
	msg.Type = message.GetUnreadMsgInfoType
	var getInfo message.GetUnreadMsgInfo
	getInfo.UserId = CurrentUserObj.UserId
	data, err := json.Marshal(getInfo)
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
func (this *SmsProcess) ShowUnreadMsgInfo(msg message.Message) {
	var returnInfo message.UnreadMsgInfoReturn
	err := json.Unmarshal([]byte(msg.Data), &returnInfo)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(msg.Data), returnInfo) err", err)
		return
	}
	for index, val := range returnInfo.UnreadMsgInfo {
		fmt.Println("发信人：", strings.Replace(index, CurrentUserObj.User.UserId+"_", "", -1), "发信数量：", val)
	}
	return
}

func (this *SmsProcess) GetUnreadMsgDetail(fromUserId string) {
	var msg message.Message
	msg.Type = message.GetUnreadMsgType
	var getInfo message.GetUnreadMsg
	getInfo.UserId = CurrentUserObj.UserId
	getInfo.FromUserId = fromUserId
	data, err := json.Marshal(getInfo)
	if err != nil {
		fmt.Println("getInfo json.Marshal err", err)
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

func (this *SmsProcess) ReadUnreadMsgDetail(msg message.Message) {
	var returnInfo message.UnreadMsgReturn
	err := json.Unmarshal([]byte(msg.Data), &returnInfo)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(msg.Data), returnInfo) err", err)
		return
	}
	// var res message.UserSmsMsg
	for index, val := range returnInfo.Data {
		// json.Unmarshal(val, &res)
		fmt.Println("第", index+1, "条消息，内容：", val.Content)
	}
	return
}
