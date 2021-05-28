package processes

import (
	"fmt"
	"go_code/chatPro/client/utils"
	"net"
	"os"
)

type Server struct{
}

func ShowMenu() {
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
				fmt.Println("show online list")
			case 2:
				fmt.Println("send message")
			case 3:
				fmt.Println("message list")
			case 4:
				fmt.Println("exit the system")
				os.Exit(0)
			default:
				fmt.Println("scanf number 1 - 4")
		}
}

func processServerMsg(conn net.Conn) {
	// start a new pkg transfer instance and read server message
	tf := &utils.Transfer{
		Conn : conn,
	}
	for {
		fmt.Println("client is reading msg from server")
		msg, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg wrong,2 error=", err)
			return
		}
		fmt.Printf("mes=%v\n", msg)
	}
}