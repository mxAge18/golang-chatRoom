package processes

import (
	"chatPro/common/message"
	"chatPro/server/model"
	"chatPro/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type SmsServerProcess struct {
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
	for id, up := range ServerUserManger.onlineUsers {
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
	val, ok := ServerUserManger.onlineUsers[smsSingleMsg.To]
	if ok {
		this.SendMsgToEachOnlineUser(data, val.Conn)
	} else {
		conn := model.MyUserDBO.Pool.Get()
		defer conn.Close()
		_, err = model.MyUserDBO.GetByFiledUserId(conn, smsSingleMsg.To)
		if err != nil {
			fmt.Println("the userid is falut, can't send message to him")
		} else {
			// store the message
			fmt.Println("the userid is not online, when he online will receive this")
			err = model.ThisUserMsgDao.StoreUnreadMsg(smsSingleMsg)
			if err != nil {
				fmt.Println("message send success")
			}
		}

	}
}

func (this *SmsServerProcess) SendMsgInfoToSomeOne(msg *message.Message) {
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
	val, ok := ServerUserManger.onlineUsers[smsSingleMsg.To]
	if ok {
		this.SendMsgToEachOnlineUser(data, val.Conn)
	} else {
		conn := model.MyUserDBO.Pool.Get()
		defer conn.Close()
		_, err = model.MyUserDBO.GetByFiledUserId(conn, smsSingleMsg.To)
		if err != nil {
			fmt.Println("the userid is falut, can't send message to him")
		} else {
			// store the message
			fmt.Println("the userid is not online, when he online will receive this")
			err = model.ThisUserMsgDao.StoreUnreadMsg(smsSingleMsg)
			if err != nil {
				fmt.Println("message send success")
			}
		}

	}
}

func (this *SmsServerProcess) SendUnreadMsgInfoToSomeOne(msg *message.Message) {

	var getUnreadMsgInfo message.GetUnreadMsgInfo
	json.Unmarshal([]byte(msg.Data), &getUnreadMsgInfo)
	conn := model.MyUserDBO.Pool.Get()
	defer conn.Close()
	_, err := model.MyUserDBO.GetByFiledUserId(conn, getUnreadMsgInfo.UserId)
	if err != nil {
		fmt.Println("the userid is falut, can't send message to him")
	} else {
		// get the message info and send to the user
		fmt.Println("未读消息如下")
		data := model.ThisUserMsgDao.GetUnreadMsgInfo(getUnreadMsgInfo.UserId)
		res, err := json.Marshal(data)
		if err != nil {
			fmt.Println("message send success")
		}

		var returnMsg message.Message
		returnMsg.Type = message.GetUnreadMsgInfoReturnType
		returnMsg.Data = string(res)
		result, err := json.Marshal(returnMsg)
		if err != nil {
			fmt.Println("msg json.Marshal err", err)
		}
		val, ok := ServerUserManger.onlineUsers[getUnreadMsgInfo.UserId]
		if ok {
			this.SendMsgToEachOnlineUser(result, val.Conn)
		}
	}
}

func (this *SmsServerProcess) SendUnreadMsgDetailToSomeOne(msg *message.Message) {

	var getInfo message.GetUnreadMsg
	json.Unmarshal([]byte(msg.Data), &getInfo)
	conn := model.MyUserDBO.Pool.Get()
	defer conn.Close()
	_, err := model.MyUserDBO.GetByFiledUserId(conn, getInfo.FromUserId)
	if err != nil {
		fmt.Println("the userid is falut")
	} else {
		// get the message info and send to the user
		fmt.Println("未读消息如下")
		fmt.Println("getInfo.UserId",getInfo.FromUserId)
		fmt.Println("getInfo.FromUserId",getInfo.FromUserId)
		data := model.ThisUserMsgDao.GetUnreadMsgDetail(getInfo.UserId, getInfo.FromUserId)
		res, err := json.Marshal(data)
		if err != nil {
			fmt.Println("消息获取错误")
		}

		var returnMsg message.Message
		returnMsg.Type = message.UnreadMsgReturnType
		returnMsg.Data = string(res)
		result, err := json.Marshal(returnMsg)
		if err != nil {
			fmt.Println("msg json.Marshal err", err)
		}
		val, ok := ServerUserManger.onlineUsers[getInfo.UserId]
		if ok {
			this.SendMsgToEachOnlineUser(result, val.Conn)
		}
	}
}
