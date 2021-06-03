package main

import (
	"fmt"
	"go_code/chatPro/server/model"
	"net"
	"time"
)
func process(conn net.Conn) {
	defer conn.Close()
	subProcessor := &Processor{
		Conn : conn,
	}
	err := subProcessor.SubProcessor()
	if err != nil {
		fmt.Println("客户端和服务器通讯协程错误=err", err)
		return 
	}

}
func initUserDBO() {
	model.MyUserDBO = model.NewUserDBO(pool)
}

func init() {
	initPool("0.0.0.0:6379", 16, 0, 300 * time.Second) // 服务器启动初始化连接池
	initUserDBO()
}
func main() {
	fmt.Println("start listen on the port 8888")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	defer listen.Close()
	if err != nil {
		fmt.Println("listen error is", err)
		return
	}
	for {
		fmt.Println("start accpet the connection")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept() error is", err)
			// return
		}
		go process(conn)
	}

}
