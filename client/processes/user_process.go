package processes

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/chatPro/client/utils"
	"go_code/chatPro/common/message"
	"net"
	"os"
)

type UserProcess struct {
	//don't needs at this time
}

func (this *UserProcess) Login(userId string, userPwd string) (err error) {
	// user login task
	conn, err := net.Dial("tcp", "192.168.1.106:8888")
	if err != nil {
		fmt.Println("net connection error", err)
		return
	}
	defer conn.Close()

	var msg message.Message
	msg.Type = message.LoginMsgType
	var loginMsg message.LoginMsg
	loginMsg.UserId = userId
	loginMsg.UserPwd = userPwd
	// 序列化要传输的数据
	data, err := json.Marshal(loginMsg)
	if err != nil {
		fmt.Println("data json marshal error", err)
		return
	}
	msg.Data = string(data)

	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("msg json marshal error", err)
		return
	}
	var pkglen uint32
	pkglen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkglen)
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("write data len wrong, err=", err)
		return
	}
	fmt.Printf("data len send to server %d, content is %s\n", len(data), string(data))

	_, err = conn.Write(data)
	if nil != err {
		fmt.Println("write data wrong, err=", err)
		return
	}

	tr := &utils.Transfer{
		Conn: conn,
	}
	msg, err = tr.ReadPkg()
	if nil != err {
		fmt.Println("readPkg(conn) wrong1, err=", err)
		return
	}
	var loginResMsg message.LoginResultMsg
	err = json.Unmarshal([]byte(msg.Data), &loginResMsg)
	if nil != err {
		fmt.Println(" json.Unmarshal([]byte(msg.Data), &loginResMsg) wrong, err=", err)
		return
	}
	if loginResMsg.Code == 200 {
		fmt.Println("client received the server response, login successful")

		// show online users
		for _, v := range(loginResMsg.Data) {
			fmt.Println("online user:", v)
		}
		// start a new process connect with the server
		// sP := &Server{}
		go processServerMsg(conn)
		// start show menu
		for {
			ShowMenu()
		}
	} else {
		fmt.Println(loginResMsg.Error)
	}
	return
}

func (this *UserProcess) Register(userId, userName, UserPwd string) (err error) {
	conn, err := net.Dial("tcp", "192.168.1.106:8888")
	if err != nil {
		fmt.Println("net connection error", err)
		return
	}
	defer conn.Close()

	var msg message.Message
	msg.Type = message.RegisterMsgType
	var registerMsg message.RegisterMsg
	registerMsg.User.UserId = userId
	registerMsg.User.UserPwd = UserPwd
	registerMsg.User.UserName = userName
	data, err := json.Marshal(registerMsg)
	if err != nil {
		fmt.Println("data json marshal error", err)
		return
	}
	msg.Data = string(data)
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("msg json marshal error", err)
		return
	}

	tr := &utils.Transfer{
		Conn: conn,
	}
	err = tr.WritePkg(data)
	if err != nil {
		fmt.Println("tr.WritePkg(data) error", err)
		return
	}
	msg, err = tr.ReadPkg()
	if nil != err {
		fmt.Println("register readPkg(conn) wrong, err=", err)
		return
	}
	var registerResultMsg message.RegisterResultMsg
	err = json.Unmarshal([]byte(msg.Data), &registerResultMsg)
	if nil != err {
		fmt.Println(" json.Unmarshal([]byte(msg.Data), &RegisterResultMsg) wrong, err=", err)
		return
	}
	if registerResultMsg.Code == 200 {
		fmt.Println("register successful")
	} else {
		fmt.Println(registerResultMsg.Error)
	}
	os.Exit(0)
	return
}
