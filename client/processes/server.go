package processes

import (
	"chatPro/client/utils"
	"chatPro/common/message"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type Server struct {
	Conn         net.Conn
	ShowMainMenu bool
}

func (this *Server) ShowMenu() {
	sp := SmsProcess{}
	for {
		fmt.Println("this.ShowMainMenu,", this.ShowMainMenu)
		fmt.Println("--welcome to chat room------------")
		fmt.Println("--1 user online list--------------")
		fmt.Println("--2 send msg to all online users--")
		fmt.Println("--3 unread message----------------")
		if !this.ShowMainMenu {
			fmt.Println("--------3-1 please input name to receive msg----")
		}
		fmt.Println("--4 send message to someone-------")
		fmt.Println("--5 exit the system---------------")
		fmt.Println("--please choose(1-5)--------------")

		var key string
		fmt.Scanf("%s\n", &key)
		switch key {
		case "1":
			ClientUserMangerObj.OutputOnlineUsers()
		case "2":
			var msg string
			fmt.Println("please input message")
			fmt.Scanln(&msg)
			sp.SendGroup(msg)
		case "3":
			fmt.Println("message list")
			sp.GetUnreadMsg()
		case "4":
			ClientUserMangerObj.OutputOnlineUsers()
			fmt.Println("please type the userId of the person you want to send message")
			var userId string
			fmt.Scanln(&userId)
			var msg string
			fmt.Println("please input message")
			fmt.Scanln(&msg)
			sp.SendSingleMsg(msg, userId)
		case "5":
			fmt.Println("exit the system")
			os.Exit(0)
		case "3-1":
			if !this.ShowMainMenu {
				fmt.Println("----please input name to receive msg----")
				var inputName string
				fmt.Scanf("%s\n", &inputName)
				sp.GetUnreadMsgDetail(inputName)
			}
		default:
			fmt.Println("scanf number 1 - 5")
		}
	}
}

func (this *Server) ProcessServerMsg() {
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
				// ??????????????????????????????
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

			case message.SmsMsgSingleReturnType:
				sp := &SmsProcess{}
				sp.ReadSingleMsg(msg)
			case message.GetUnreadMsgInfoReturnType:
				sp := &SmsProcess{}
				sp.ShowUnreadMsgInfo(msg)
				this.ShowMainMenu = false
			case message.UnreadMsgReturnType:
				sp := &SmsProcess{}
				sp.ReadUnreadMsgDetail(msg)
			default:
				fmt.Printf("msg????????????=%v\n", msg)
		}

	}

}
