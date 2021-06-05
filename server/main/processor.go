package main

import (
	"fmt"
	"go_code/chatPro/common/message"
	"go_code/chatPro/server/processes"
	"go_code/chatPro/server/utils"
	"io"
	"net"
)
type Processor struct {
	Conn net.Conn
}
func (this *Processor) serverProcessMsg(msg *message.Message) (err error) {
	switch msg.Type {
		case message.LoginMsgType:
			userPro := &processes.UserProcessor{
				Conn : this.Conn,
			}
			err = userPro.ServerProcessLogin(msg)
		case message.RegisterMsgType:
			//deal with the register
			userPro := &processes.UserProcessor{
				Conn : this.Conn,
			}
			err = userPro.ServerProcessRegister(msg)
		case message.SmsMsgType:
			smsPro := &processes.SmsServerProcess{}
			smsPro.SendGroupMsg(msg)
		default:
			fmt.Println("message type is not right", msg.Type)

	}
	return

}
func (this *Processor) SubProcessor() (err error) {
	for {
		tr := &utils.Transfer{
			Conn: this.Conn,
		}
		msg, err := tr.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("quit the client so server stop listion client msg")
			} else {
				fmt.Println("readPkg error, ", err)
			}
			return err
		}
		fmt.Println("msg is", msg)

		err = this.serverProcessMsg(&msg)
		if err != nil {
			return err
		}
		// important:原来直接写return err,客户端协程监听服务端发送的消息就收到EOF的错误。改成 if err != nil :return err
	}
}


