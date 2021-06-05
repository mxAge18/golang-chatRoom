package processes

import (
	"encoding/json"
	"fmt"
	"go_code/chatPro/client/utils"
	"go_code/chatPro/common/message"
	"net"
	"os"
)
type Server struct{
	Conn net.Conn
}
func(this *Server) ShowMenu() {
	fmt.Println("------------welcome to chat room------------")
	fmt.Println("------------1 user online list--------------")
	fmt.Println("------------2 send message------------------")
	fmt.Println("------------3 message list------------------")
	fmt.Println("------------4 exit the system---------------")
	fmt.Println("------------please choose(1-4)--------------")
	var key int
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		ClientUserMangerObj.OutputOnlineUsers()
	case 2:
		sp:= SmsProcess{}
		var msg string
		fmt.Println("please input message")
		fmt.Scanln(&msg)
		sp.SendGroup(msg)
	case 3:
		fmt.Println("message list")
	case 4:
		fmt.Println("exit the system")
		os.Exit(0)
	default:
		fmt.Println("scanf number 1 - 4")
	}
}

func(this *Server) ProcessServerMsg() {
	// start a new pkg transfer instance and read server message
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	for {
		msg, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg wrong,2 error=", err)
			return
		}
		switch msg.Type {
		case message.NotifyUserStatusMsgType:
			var NotifyUserStatusMsg message.NotifyUserStatusMsg
			json.Unmarshal([]byte(msg.Data), &NotifyUserStatusMsg)
			// 更新到用户在线信息中
			user := &message.User{
				UserId:     NotifyUserStatusMsg.UserId,
				UserName:   NotifyUserStatusMsg.UserId,
				UserStatus: NotifyUserStatusMsg.UserStatus,
			}
			ClientUserMangerObj.AddOnlineUser(user)
			ClientUserMangerObj.OutputOnlineUsers()
		case message.GroupReturnMsgType:
			sp := &SmsProcess{}
			sp.ReadGroupMsg(msg)
		default:
			fmt.Printf("msg无法解析=%v\n", msg)
		}

	}

}
